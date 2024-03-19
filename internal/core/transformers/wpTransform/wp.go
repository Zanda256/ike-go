package wpTransform

import (
	"regexp"

	"github.com/Zanda256/ike-go/internal/core/importers/wpImport/stores/wpImportDb"
	"github.com/google/uuid"
)

const (
	textEmbedding3Small    = "text-embedding-3-small"
	OpenAIMaxContext       = 8191
	MarkdownFormat         = "md"
	MinChunkSize           = 212
	cl100k_baseEncoding    = "cl100k_base"
	TogetherComputerSubStr = "togethercomputer/"
	TextEmbeddingSubStr    = "text-embedding"
)

var (
	mdLinkPattern             = regexp.MustCompile("\\[[^\\]]*\\]\\([^\\)]*\\)")
	TogetherAPIKey            string
	OpenAPIKey                string
	OpenAPIEmbeddingEndpoint  = "https://api.openai.com/v1/embeddings"
	TogetherEmbeddingEndPoint = "https://api.together.xyz/v1/embeddings"
)

type Storer interface {
	GetSourcesByHosts(hosts []string) ([]wpImportDb.Source, error)
	GetDownloadBySource(sourceID uuid.UUID) (wpImportDb.Download, error)
}

type WpTransformer struct {
	SourceManager *SourceTransformManager
}

func (wpT *WpTransformer) Transform(hosts []string) error {
	srcs, err := wpT.SourceManager.Store.GetSourcesByHosts(hosts)
	if err != nil {
		return err
	}
	_ = srcs
	return err
}

// 1. Connect to db

// 2.  process_sources
// 2.1 get sources from db where host = given hosts
// 2.2 process the each source in order up to the {{limit}} given

// 3.  process_source
// 3.1 find out what this is for embedding_info = {"name": "text-embedding-3-small", "max_context": 8191} :- These OpenAi API config for generating embeddings for our text.
// 3.2 get download for the source_id from db

// 4.  processing each download into a document
// 4.1 Deserialize the body of the download into json
// 4.2 Get the contents of body["content"]["rendered"] and convert them from html into markdown
// 4.3 Change the contents of  from string to time.Time in this format "%Y-%m-%dT%H:%M:%S"
// 4.4 Do the same for body["date_gmt"]

// 4.5 Detect language of the contents you got from body["content"]["rendered"]
// 4.6 Build the document of this structure
//   document = {
//      "modified_at": parse(modified_gmt.isoformat()),
//      "published_at": parse(date_gmt.isoformat()),
//      "format": "md",
// 	 }

// return document, content, natural_lang

// Add this info to the document struct/map
// document["max_chunk_size"] = embedding_info["max_context"]
// document["min_chunk_size"] = 212

// 5. process_metadata
// body = json.loads(download[6])
// metadata = {
// 	"document_title": md(body["title"]["rendered"]),
// 	"document_description": md(body["excerpt"]["rendered"]),
// 	"links_count": number_of_links(content),
// 	"canonical_url": body["link"],
// }
