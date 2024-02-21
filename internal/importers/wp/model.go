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
