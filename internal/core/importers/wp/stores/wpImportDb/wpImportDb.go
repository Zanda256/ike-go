package wpImportDb

import (
	"context"
	"fmt"

	"github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/google/uuid"
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

func (st *Store) InsertSource(s Source) (uuid.UUID, error) {
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

func (st *Store) InsertDownload(d Download) (uuid.UUID, error) {
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
	_, err = tx.Prepare(ctx, "insert_download", "INSERT INTO downloads (id, source_id, created_at, updated_at, status_code, headers, body) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return uuid.Nil, err
	}

	// Execute the prepared statement with values from the Download struct
	_, err = tx.Exec(ctx,
		"insert_download",
		d.ID,
		d.SourceID,
		d.CreatedAt,
		d.UpdatedAt,
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

//docker run -it --rm --network some-network postgres psql -h some-postgres -U postgres

//docker exec -it <container_name> sh
//Fetch and follow the logs of a container:
//docker logs -f <container_name>
//To inspect a running container:
//docker inspect <container_name> (or <container_id>)
//
//Create and run a container from an image, with a custom name:
//docker run --name <container_name> <image_name>
//Run a container with and publish a container’s port(s) to the host.
//docker run -p <host_port>:<container_port> <image_name>
//Run a container in the background
//docker run -d <image_name>
docker volume create postgres-volume
docker run --name ike-postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres --volume postgres-volume:/var/lib/postgresql/data  -d postgres:latest
