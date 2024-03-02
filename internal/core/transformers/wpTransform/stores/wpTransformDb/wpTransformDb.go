package wpTransformDb

import (
	"context"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport/stores/wpImportDb"
	"github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
)

type Store struct {
	log *logger.Logger
	db  *dbsql.DB
}

func NewStore(log *logger.Logger, db *dbsql.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

func (s *Store) GetSourcesByHosts(hosts []string) ([]wpImportDb.Source, error) {
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

	// Prepare the SQL statement
	//_, err = tx.Prepare(ctx, "insert_source", "INSERT INTO sources (id, author_email, raw_url, scheme, host, path, query, active_domain,\"format\", created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	//if err != nil {
	//	return uuid.Nil, err
	//}
	//
	//_, err = tx.Exec(ctx,
	//	"insert_source",
	//	s.ID,
	//	s.AuthorEmail,
	//	s.RawURL,
	//	s.Scheme,
	//	s.Host,
	//	s.Path,
	//	s.Query,
	//	s.ActiveDomain,
	//	s.Format,
	//	s.CreatedAt,
	//	s.UpdatedAt,
	//)
	//if err != nil {
	//	return uuid.Nil, err
	//}
	//return s.ID, nil
}
