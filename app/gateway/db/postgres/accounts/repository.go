package accounts

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
)

type Repository struct {
	db postgres.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
