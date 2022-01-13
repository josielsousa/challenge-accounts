package accounts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
)

func (r *Repository) Insert(ctx context.Context, trf transfers.Transfer) error {
	const op = `Repository.Transfers.Insert`
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
            account_origin_id,
            account_destination_id,
            amount,
            created_at,
            updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at
    `

	row := r.db.QueryRow(
		ctx,
		query,
		trf.ID,
		trf.AccountOriginID,
		trf.AccountDestinationID,
		trf.Amount,
		trf.CreatedAt,
		trf.UpdatedAt,
	)

	err := row.Scan(
		&trf.ID,
		&trf.CreatedAt,
		&trf.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return fmt.Errorf("%s-> %s: %w", op, "on insert transfer", accounts.ErrAccountAlreadyExists)
		}

		return fmt.Errorf("%s-> %s: %w", op, "on insert transfer", err)
	}

	return nil
}
