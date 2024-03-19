package wpTransformDb

import (
	"context"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport/stores/wpImportDb"
	"github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/google/uuid"
)

type DownloadStore struct {
	log *logger.Logger
	db  *dbsql.DB
}

func NewStore(log *logger.Logger, db *dbsql.DB) *DownloadStore {
	return &DownloadStore{
		log: log,
		db:  db,
	}
}

func (s *DownloadStore) GetSourcesByHosts(hosts []string) ([]wpImportDb.Source, error) {
	ctx := context.Background()
	// Begin a new transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	// TODO: Add condition to retrieve all sources if hosts is an empty array

	var res []wpImportDb.Source
	rows := tx.QueryRow(ctx, `SELECT * FROM sources WHERE host = ANY($1)`, hosts)
	if err = rows.Scan(res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *DownloadStore) GetDownloadBySource(sourceID uuid.UUID) (wpImportDb.Download, error) {
	ctx := context.Background()
	// Begin a new transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return wpImportDb.Download{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var res wpImportDb.Download
	row := tx.QueryRow(ctx, `SELECT * FROM downloads WHERE source_id = $1`, sourceID)
	if err = row.Scan(res); err != nil {
		return wpImportDb.Download{}, err
	}

	return res, nil
}
