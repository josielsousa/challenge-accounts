package transfers

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/accounts"
)

type Repository struct {
	db      *pgxpool.Pool
	accRepo accounts.Repository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:      db,
		accRepo: *accounts.NewRepository(db),
	}
}
