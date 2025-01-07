package transfers

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

func (r *Repository) Insert(ctx context.Context, trf transfers.TransferData) error {
	const op = `Repository.Transfers.Insert`

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on open transaction", err)
	}

	//nolint:errcheck
	defer tx.Rollback(ctx)

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
            account_origin_id,
            account_destination_id,
            amount,
            created_at,
            updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at
    `

	row := tx.QueryRow(
		ctx,
		query,
		trf.ID,
		trf.AccountOriginID,
		trf.AccountDestinationID,
		trf.Amount,
		trf.CreatedAt,
		trf.UpdatedAt,
	)

	err = row.Scan(
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

	err = r.accRepo.UpdateBalance(ctx, tx, trf.AccountOriginID, trf.AccountOrigin.Balance)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on update account origin", err)
	}

	err = r.accRepo.UpdateBalance(ctx, tx, trf.AccountDestinationID, trf.AccountDestination.Balance)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on update account destination", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on commit transaction", err)
	}

	return nil
}
