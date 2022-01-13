package pgtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
)

func TransfersInsert(t *testing.T, db *pgxpool.Pool, trf transfers.Transfer) error {
	t.Helper()

	const op = `PgTest.TransfersInsert`
	if len(trf.ID) <= 0 {
		trf.ID = uuid.NewString()
	}

	if trf.CreatedAt.IsZero() {
		trf.CreatedAt = time.Now().In(time.Local)
	}

	if trf.UpdatedAt.IsZero() {
		trf.UpdatedAt = time.Now().In(time.Local)
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

func GetTransfer(t *testing.T, db *pgxpool.Pool, id string) (transfers.Transfer, error) {
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

	var trf transfers.Transfer

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
		return transfers.Transfer{}, fmt.Errorf("%s-> %s:%w", op, "on query transfer", err)
	}

	return trf, nil
}
