package wpImport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport/stores/wpImportDb"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/Zanda256/ike-go/pkg-foundation/web"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// Import url - https://wsform.com/wp-json/wp/v2/knowledgebase

type Storer interface {
	InsertSource(s wpImportDb.Source) (uuid.UUID, error)
	InsertDownload(d wpImportDb.Download) (uuid.UUID, error)
}

type ImportManager struct {
	Storage   Storer
	WebClient *web.ClientProvider
	Log       *logger.Logger
}

func NewWordPressImporter(log *logger.Logger, client *web.ClientProvider, store Storer) *ImportManager {
	return &ImportManager{
		Storage:   store,
		WebClient: client,
		Log:       log,
	}
}

func (wpi *ImportManager) Import(fullURL string) error { // errChan chan error
	resultsPerPage := 100
	page := 1
	fmt.Printf("url to import: %s", fullURL)
	//wpi.Log.Info(context.Background(), "url to import: ", fullURL)
	for {
		url := fmt.Sprintf("%s?page=%d&per_page=%d", fullURL, page, resultsPerPage)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		// make headers map
		headers := make(map[string]string)
		headers["User-Agent"] = "ike-script"
		headers["Accept"] = "application/json"

		// Assign headers to the request
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		//if wpi.Storage != nil {
		//	fmt.Printf("\nstorage pointer is not nil\n")
		//}
		//if wpi.WebClient == nil {
		//	if wpi.Log == nil {
		//		fmt.Printf("\nlog and http pointers are nil\n")
		//	}
		//	fmt.Printf("\nlog and http pointers are nil\n")
		//	//wpi.Log.Info(context.Background(), "nil http client")
		//	return errors.New("nil http client")
		//}
		resp, err := wpi.WebClient.SendRequest(req)
		if err != nil {
			wpi.Log.Error(context.Background(), "error encountered:", err.Error())
			return err
		}
		if resp.StatusCode == 400 {
			break
		}
		var apiResults []map[string]any
		err = json.Unmarshal(resp.Body, &apiResults)
		if err != nil {
			wpi.Log.Error(context.Background(), "error encountered:", err.Error())
			return err
		}
		//fmt.Printf("\napiResults : %+v\n", apiResults)
		for _, result := range apiResults {
			var id float64
			idAny, ok := result["id"]
			if ok {
				var cast bool
				if id, cast = idAny.(float64); !cast {
					wpi.Log.Warn(context.Background(), "failed to cast id, skipping record", result["id"])
					continue
				}
			} else {
				wpi.Log.Warn(context.Background(), "no id found, skipping record")
				continue
			}
			postURL := fmt.Sprintf("%s/%d", fullURL, int(id))
			fmt.Printf("\npostURL : %+v\n", postURL)
			s, err := wpi.fetchAndProcessPost(postURL)
			if err != nil {
				wpi.Log.Error(context.Background(), "error encountered:", err.Error())
				return err
			}
			wpi.Log.Info(context.Background(), "success for: ", s)
		}
		page += 1
	}
	return nil
}

func (wpi *ImportManager) fetchAndProcessPost(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	attemptedAt := time.Now()
	res, err := wpi.WebClient.SendRequest(req)
	if err != nil {
		wpi.Log.Error(context.Background(), "error encountered:", err.Error())
		return "", err
	}
	downloadedAt := time.Now()

	// convert URL to source and save it to db
	source, err := wpImportDb.ToSource(url)
	if err != nil {
		wpi.Log.Error(context.Background(), err.Error())
		return "", err
	}
	source.CreatedAt = time.Now()
	source.UpdatedAt = time.Now()
	sid, err := wpi.Storage.InsertSource(source)
	if err != nil {
		wpi.Log.Error(context.Background(), err.Error())
		return "", err
	}

	// build the download struct and save it to db
	download := wpImportDb.ToDownload(res, source.ID)
	download.AttemptedAt = attemptedAt
	download.DownloadedAt = downloadedAt
	did, err := wpi.Storage.InsertDownload(download)
	if err != nil {
		wpi.Log.Error(context.Background(), err.Error())
		return "", err
	}
	return fmt.Sprintf("source %s - download %s", sid.String(), did.String()), nil
}
