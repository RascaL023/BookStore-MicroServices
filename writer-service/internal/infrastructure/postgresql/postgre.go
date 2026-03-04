package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgrePool(uri string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	
	pool, err := pgxpool.New(ctx, uri)
	if err != nil { return nil, err }

	return pool, nil
}
