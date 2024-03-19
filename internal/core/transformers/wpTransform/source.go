package wpTransform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport/stores/wpImportDb"
	"github.com/Zanda256/ike-go/internal/core/transformers/wpTransform/stores/wpTransformDb"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/Zanda256/ike-go/pkg-foundation/web"
	"github.com/google/uuid"
	lingua "github.com/pemistahl/lingua-go"
	"github.com/pkoukk/tiktoken-go"
	"log"
	"net/http"
	"strings"
	"time"
)

type SourceTransformManager struct {
	httpClient web.ClientProvider
	log        logger.Logger
	Store      Storer
}

func (sm *SourceTransformManager) processSource(src wpImportDb.Source) {
	openAIcfg := struct {
		name       string
		maxContext int
	}{
		name:       textEmbedding3Small,
		maxContext: OpenAIMaxContext,
	}
	download, err := sm.Store.GetDownloadBySource(src.ID)
	if err != nil {
		sm.log.Error(context.Background(), "error retrieving download: ", err.Error())
		return
	}
	body := make(map[string]any)
	err = json.Unmarshal(download.Body, &body)
	if err != nil {
		sm.log.Error(context.Background(), "error Unmarshalling download bogy: ", err.Error())
		return
	}

	// Process document - take in body -> return error
	var renderedContentHTML string
	//content := make(map[string]any)
	if content, ok := body["content"].(map[string]any); ok {
		if rendered, ok := content["rendered"].(string); ok {
			renderedContentHTML = rendered
		}
	}

	converter := md.NewConverter("", true, nil)
	renderedContentMd, err := converter.ConvertString(renderedContentHTML)
	if err != nil {
		log.Fatal(err)
	}
	var modifiedGmtTime time.Time
	var dateGmtTime time.Time
	modifiedGmtStr, ok := body["modified_gmt"].(string)
	if ok { // "%Y-%m-%dT%H:%M:%S"
		modifiedGmtTime, err = time.Parse(time.DateTime, modifiedGmtStr)
		if err != nil {
			log.Fatal(err)
		}
	}
	dateGmtStr, ok := body["date_gmt"].(string)
	if ok { // "%Y-%m-%dT%H:%M:%S"
		dateGmtTime, err = time.Parse(time.DateTime, dateGmtStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	// detect language of body["content"]["rendered"]
	languages := lingua.AllLanguages()
	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()
	var naturalLang string
	if language, exists := detector.DetectLanguageOf(renderedContentMd); exists {
		naturalLang = language.String()
		fmt.Println(language.String())
	}
	doc := wpTransformDb.Document{
		ID:           uuid.UUID{},
		SourceID:     uuid.UUID{},
		DownloadID:   uuid.UUID{},
		Format:       MarkdownFormat,
		IndexedAt:    time.Time{},
		MinChunkSize: MinChunkSize,
		MaxChunkSize: openAIcfg.maxContext,
		PublishedAt:  dateGmtTime,
		ModifiedAt:   modifiedGmtTime,
		WpVersion:    "",
	}
	fmt.Printf("\nNatural Lang : %+v\n", naturalLang)
	fmt.Printf("\ndoc : %+v\n", doc)
	// document processed successfully

	// process metadata
	//meta := wpTransformDb.DocumentMeta{
	//	ID:         uuid.UUID{},
	//	DocumentID: uuid.UUID{},
	//	Key:        "",
	//	Meta:       nil,
	//	CreatedAt:  time.Time{},
	//}

	docTitle := ""
	if title, ok := body["title"].(map[string]string); ok {
		if rendered, ok := title["rendered"]; ok {
			docTitle = rendered
		}
	}

	docTitleMd, err := converter.ConvertString("<p>" + docTitle + "</p>")
	if err != nil {
		log.Fatal(err)
	}

	docDesc := ""
	if excerpt, ok := body["excerpt"].(map[string]any); ok {
		if rendered, ok := excerpt["rendered"].(string); ok {
			docDesc = rendered
		}
	}

	docDescMd, err := converter.ConvertString(docDesc)
	if err != nil {
		log.Fatal(err)
	}
	canonicalURL := body["link"].(string)
	var metaData = map[string]any{
		"documentTitle":       docTitleMd,
		"documentDescription": docDescMd,
		"linksCount":          numLinks(renderedContentMd),
		"canonicalURL":        canonicalURL,
	}
	// process metadata success
	fmt.Printf("\nmetaData : %+v\n", metaData)

	//processing chunks

	// if you don't want download dictionary at runtime, you can use offline loader
	// tiktoken.SetBpeLoader(tiktoken_loader.NewOfflineLoader())
	tke, err := tiktoken.GetEncoding(cl100k_baseEncoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	// encode
	tokenizedContent := tke.Encode(renderedContentMd, nil, nil)

	//tokens
	fmt.Println(tokenizedContent)
	numTokens := len(tokenizedContent)
	fmt.Println(numTokens)

	if numTokens > openAIcfg.maxContext {
		chunkSize := openAIcfg.maxContext
		chunks := make([][]int, 0)
		for i := 0; i < numTokens; i += chunkSize {
			chunk := tokenizedContent[i : i+chunkSize]
			chunks = append(chunks, chunk)
		}
		chunkObjs := make([]map[string]any, 0)
		for _, chunkTokens := range chunks {
			decodedToken := tke.Decode(chunkTokens)
			err = sm.EmbedDoc(decodedToken)
			if err != nil {
				err = fmt.Errorf("getEncoding: %v", err)
				return
			}
		}
		fmt.Printf("\nchunk_objs : %+v\n", chunkObjs)

	}

}

func (sm *SourceTransformManager) EmbedDoc(token string) error {
	model := textEmbedding3Small
	bodyMap := map[string]any{
		"input": strings.Replace(token, "\n", " ", -1),
		"model": model,
	}

	// make headers map
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	url := ""

	switch {
	case strings.Contains(model, TogetherComputerSubStr):
		url = TogetherEmbeddingEndPoint
		headers["Authorization"] = fmt.Sprintf("Bearer %s", TogetherAPIKey)

	case strings.Contains(model, TextEmbeddingSubStr):
		url = OpenAPIEmbeddingEndpoint
		headers["Authorization"] = fmt.Sprintf("Bearer %s", OpenAPIKey)
		bodyMap["encoding_format"] = "float"

	default:
		log.Fatalf("Unsupported embedding model: %s", model)
	}
	body, err := json.Marshal(bodyMap)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return err
	}
	b := bytes.NewBuffer(body)

	// Create a new POST request
	request, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Assign headers to the request
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err := sm.httpClient.SendRequest(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	fmt.Printf("Embed Response %+v", resp)
	return nil
}

// Build and set headers

//type embeddingInfo

func numLinks(s string) int {
	return len(mdLinkPattern.FindAllString(s, -1))
}
