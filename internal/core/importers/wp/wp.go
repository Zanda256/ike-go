package wp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	wpImportdb "github.com/Zanda256/ike-go/internal/core/importers/wp/stores/wpImportDb"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/Zanda256/ike-go/pkg-foundation/web"
	"github.com/google/uuid"
)

type Storer interface {
	InsertSource(s wpImportdb.Source) (uuid.UUID, error)
	InsertDownload(d wpImportdb.Download) (uuid.UUID, error)
}

type ImportManager struct {
	Storage   Storer
	webClient *web.ClientProvider
	log       *logger.Logger
}

func NewWordPressImporter(log *logger.Logger, client *web.ClientProvider, store Storer) *ImportManager {
	return &ImportManager{
		Storage:   store,
		webClient: client,
		log:       log,
	}
}

func (wpi *ImportManager) Import(fullURL string) error { // errChan chan error
	resultsPerPage := 100
	page := 1
	for {
		url := fmt.Sprintf("%s?page=%d&per_page=%d", fullURL, page, resultsPerPage)
		resp, err := wpi.webClient.SendRequest(http.MethodGet, url, nil)
		if err != nil {
			wpi.log.Error(context.Background(), "error encountered:", err.Error())
			return err
		}
		if resp.StatusCode == 400 {
			break
		}
		var m []map[string]any
		err = json.Unmarshal(resp.Body, m)
		if err != nil {
			wpi.log.Error(context.Background(), "error encountered:", err.Error())
			return err
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

func (wpi *ImportManager) fetchAndProcessPost(url string) (string, error) {
	res, err := wpi.webClient.SendRequest(http.MethodGet, url, nil)
	if err != nil {
		wpi.log.Error(context.Background(), "error encountered:", err.Error())
		return "", err
	}
	// convert URL to source and save it to db
	source, err := wpImportdb.ToSource(url)
	if err != nil {
		wpi.log.Error(context.Background(), err.Error())
		return "", err
	}
	sid, err := wpi.Storage.InsertSource(source)
	if err != nil {
		wpi.log.Error(context.Background(), err.Error())
		return "", err
	}
	fmt.Printf("Source ID: %s", source.ID.String())
	// build the download struct and save it to db
	download := wpImportdb.ToDownload(res, source.ID)
	did, err := wpi.Storage.InsertDownload(download)
	if err != nil {
		wpi.log.Error(context.Background(), err.Error())
		return "", err
	}
	fmt.Printf("Download ID: %s", download.ID.String())
	return fmt.Sprintf("source %s - download %s", sid.String(), did.String()), nil
}
