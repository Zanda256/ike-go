package importers

import (
	"fmt"

	db "github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg/logger"
	web "github.com/Zanda256/ike-go/pkg/web"
)

type WordPressImporter struct {
	Storage   *db.DB
	webClient *web.ClientProvider
	log       logger.Logger
}

func (wpi *WordPressImporter) Import(fullURL string, page, resultsPerPage int) {
	url := fmt.Sprintf("%s?page=%d&per_page=&d", fullURL, page, resultsPerPage)
	wpi.webClient.Sen
}
