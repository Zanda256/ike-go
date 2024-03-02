package wpTransformDb

import (
	"github.com/google/uuid"
	"time"
)

type Tag struct {
	ID         uuid.UUID
	Name       string
	created_at time.Time
}

type Chunk struct {
	ID            uuid.UUID
	DocumentID    uuid.UUID
	ParentChunkID uuid.UUID
	LeftChunkID   uuid.UUID
	RightChunkID  uuid.UUID
	Body          []byte
	ByteSize      int
	Tokenizer     string // enum?
	TokenCount    int
	NaturalLang   string // enum?
	CodeLang      string // enum?
}

type DocumentTags struct {
	ID         uuid.UUID
	DocumentID uuid.UUID
	TagID      uuid.UUID
	CreatedAt  time.Time
}

type Document struct {
	ID           uuid.UUID
	SourceID     uuid.UUID
	DownloadID   uuid.UUID
	Format       string
	IndexedAt    time.Time
	MinChunkSize int
	MaxChunkSize int
	PublishedAt  time.Time
	ModifiedAt   time.Time
	WpVersion    string
}

type DocumentMeta struct {
	ID         uuid.UUID
	DocumentID uuid.UUID
	Key        string
	Meta       []byte
	CreatedAt  time.Time
}

type Embedding struct {
	ID            uuid.UUID
	Embedding1536 string // enum? vector
	Embedding3072 string // enum? vector
	Model         string
	EmbeddedAt    time.Time
	ObjectID      uuid.UUID
	ObjectType    string
	Embedding768  string // enum? vector
}

type Request struct {
	ID           uuid.UUID
	Message      []byte
	Meta         []byte
	RequestedAt  time.Time
	ResultChunks []uuid.UUID
}
