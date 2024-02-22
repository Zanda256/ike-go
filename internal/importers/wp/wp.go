package importers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg/logger"
	"github.com/Zanda256/ike-go/pkg/web"
)

type WordPressImporter struct {
	Storage   *db.DB
	webClient *web.ClientProvider
	log       logger.Logger
}

func (wpi *WordPressImporter) Import(fullURL string, resultsPerPage int) error {
	page := 1
	for {
		url := fmt.Sprintf("%s?page=%d&per_page=&d", fullURL, page, resultsPerPage)
		resp, err := wpi.webClient.SendRequest(http.MethodGet, url, nil)
		if err != nil {
			wpi.log.Error(context.Background(), "error encountered:", err.Error())
		}
		if resp.StatusCode == 400 {
			break
		}
		var m []map[string]any
		err = json.Unmarshal(resp.Body, m)
		if err != nil {
			wpi.log.Error(context.Background(), "error encountered:", err.Error())
		}
		for _, result := range m {
			id, ok := result["id"]
			if !ok {
				wpi.log.Warn(context.Background(), "no id found, skipping record")
				continue
			}
			postURL := fmt.Sprintf("%s/%s", fullURL, id)
			wpi.fetchAndProcessPost(postURL)
		}
		page += 1
	}

}

func (wpi *WordPressImporter) fetchAndProcessPost(url string) (string, error) {
	res, err := wpi.webClient.SendRequest(http.MethodGet, url, nil)
	if err != nil {
		wpi.log.Error(context.Background(), "error encountered:", err.Error())
	}
	// convert URL to source and save it to db
	source, err := toSource(url)
	if err != nil {
		wpi.log.Error(context.Background(), err.Error())
		return "", err
	}
	fmt.Printf("Source ID: %s", source.ID.String())
	// build the download struct and save it to db
	download := toDownload(res, source.ID)
	fmt.Printf("Download ID: %s", download.ID.String())
	return fmt.Sprintf("source %s - download %s", source.ID.String(), download.ID.String()), nil
}
