package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	pgxpool.Pool
	Begin(ctx context.Context) (pgx.Tx, error)
}
