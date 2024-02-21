package importers

import (
	"time"

	"github.com/google/uuid"
)

type Source struct {
	ID           uuid.UUID
	AuthorEmail  string
	RawURL       string
	Scheme       string
	Host         string
	Path         string
	Query        string
	ActiveDomain bool
	DocFormat    string // enum
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Downloads struct {
	ID         uuid.UUID
	SourceID   uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StatusCode int
	Headers    []byte
	Body       []byte
}

type Tag struct {
	ID         uuid.UUID
	Name       string
	created_at time.Time
}
