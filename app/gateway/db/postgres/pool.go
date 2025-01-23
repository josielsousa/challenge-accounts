package postgres

import (
	"context"

	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
)

type Pool interface {
	pgxtype.Querier
	Begin(ctx context.Context) (pgx.Tx, error)
}
