package transfers

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/accounts"
)

type Repository struct {
	db      postgres.Pool
	accRepo accounts.Repository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:      db,
		accRepo: *accounts.NewRepository(db),
	}
}
