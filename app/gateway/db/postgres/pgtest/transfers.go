package pgtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
)

func TransfersInsert(t *testing.T, db *pgxpool.Pool, trf entities.Transfer) error {
	t.Helper()

	const op = `PgTest.TransfersInsert`

	if len(trf.ID) == 0 {
		trf.ID = uuid.NewString()
	}

	if trf.CreatedAt.IsZero() {
		trf.CreatedAt = time.Now().In(time.UTC)
	}

	if trf.UpdatedAt.IsZero() {
		trf.UpdatedAt = time.Now().In(time.UTC)
	}

	query := `
		INSERT INTO transfers(
			id,
			amount,
			account_origin_id,
			account_destination_id,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
    `

	row := db.QueryRow(
		context.Background(),
		query,
		trf.ID,
		trf.Amount,
		trf.AccountOriginID,
		trf.AccountDestinationID,
		trf.CreatedAt,
		trf.UpdatedAt,
	)

	err := row.Scan(
		&trf.ID,
		&trf.CreatedAt,
		&trf.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s-> %s:%w", op, "on insert transfer", err)
	}

	return nil
}

func GetTransfer(t *testing.T, db *pgxpool.Pool, id string) (entities.Transfer, error) {
	t.Helper()

	const op = `PgTest.GetTransfer`

	query := `
        SELECT
            id,
			amount,
			account_origin_id,
			account_destination_id,
			created_at,
			updated_at
		FROM transfers
		WHERE id = $1
    `

	var trf entities.Transfer

	row := db.QueryRow(
		context.Background(),
		query,
		id,
	)

	err := row.Scan(
		&trf.ID,
		&trf.Amount,
		&trf.AccountOriginID,
		&trf.AccountDestinationID,
		&trf.CreatedAt,
		&trf.UpdatedAt,
	)
	if err != nil {
		return entities.Transfer{}, fmt.Errorf("%s-> %s:%w", op, "on query transfer", err)
	}

	return trf, nil
}
