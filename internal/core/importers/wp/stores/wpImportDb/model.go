package wpImportdb

import (
	"net/url"
	"time"

	"github.com/Zanda256/ike-go/pkg/web"
	"github.com/google/uuid"
)

const JSONFormat = "json"

type Source struct {
	ID           uuid.UUID
	AuthorEmail  string
	RawURL       string
	Scheme       string
	Host         string
	Path         string
	Query        string
	ActiveDomain bool
	Format       string // enum
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ToSource(postURL string) (Source, error) {
	// Parse the URL
	parsedURL, err := url.Parse(postURL)
	if err != nil {
		//fmt.Println("Error parsing URL:", err)
		return Source{}, err
	}
	//populate source with raw_url, scheme, host, path, query, format
	s := Source{
		ID:     uuid.New(),
		RawURL: postURL,
		Scheme: parsedURL.Scheme,
		Host:   parsedURL.Host,
		Path:   parsedURL.Path,
		Query:  parsedURL.RawQuery,
		Format: JSONFormat,
	}
	return s, nil
}

type Download struct {
	ID         uuid.UUID
	SourceID   uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StatusCode int
	Headers    []byte
	Body       []byte
}

func ToDownload(raw web.Response, sourceID uuid.UUID) Download {
	return Download{
		ID:         uuid.New(),
		SourceID:   sourceID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
	}
}

type Tag struct {
	ID         uuid.UUID
	Name       string
	created_at time.Time
}
