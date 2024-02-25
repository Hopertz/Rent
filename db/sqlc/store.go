package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Store interface {
	Querier
	NewToken(id uuid.UUID, ttl time.Duration, scope string) (*TokenLoc, error)
	BulkInsert(ctx context.Context, houses []HouseBulk) error
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
