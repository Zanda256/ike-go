package wpImportDb

import (
	"context"
	"fmt"

	"github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/google/uuid"
)

type ImportStore struct {
	log *logger.Logger
	db  *dbsql.DB
}

func NewStore(log *logger.Logger, db *dbsql.DB) *ImportStore {
	return &ImportStore{
		log: log,
		db:  db,
	}
}

func (st *ImportStore) InsertSource(s Source) (uuid.UUID, error) {
	ctx := context.Background()
	// Begin a new transaction
	tx, err := st.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	// Prepare the SQL statement
	_, err = tx.Prepare(ctx, "insert_source", "INSERT INTO sources (id, author_email, raw_url, scheme, host, path, query, active_domain,\"format\", created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		return uuid.Nil, err
	}

	_, err = tx.Exec(ctx,
		"insert_source",
		s.ID,
		s.AuthorEmail,
		s.RawURL,
		s.Scheme,
		s.Host,
		s.Path,
		s.Query,
		s.ActiveDomain,
		s.Format,
		s.CreatedAt,
		s.UpdatedAt,
	)
	if err != nil {
		return uuid.Nil, err
	}
	return s.ID, nil
}

func (st *ImportStore) InsertDownload(d Download) (uuid.UUID, error) {
	ctx := context.Background()
	// Begin a new transaction
	tx, err := st.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	// Prepare the SQL statement
	_, err = tx.Prepare(ctx, "insert_download", "INSERT INTO downloads (id, source_id, attempted_at, downloaded_at, status_code, headers, body) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return uuid.Nil, err
	}

	// Execute the prepared statement with values from the Download struct
	_, err = tx.Exec(ctx,
		"insert_download",
		d.ID,
		d.SourceID,
		d.AttemptedAt,
		d.DownloadedAt,
		d.StatusCode,
		d.Headers,
		d.Body,
	)
	if err != nil {
		fmt.Println("Error executing statement:", err)
		return uuid.Nil, err
	}

	fmt.Println("Download inserted successfully.")
	return d.ID, nil
}
