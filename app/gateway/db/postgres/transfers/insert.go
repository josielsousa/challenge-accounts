package transfers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

func (r *Repository) Insert(
	ctx context.Context, trf entities.TransferData,
) error {
	const op = `Repository.Transfers.Insert`

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s-> on open transaction: %w", op, err)
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
		if errors.As(err, &pgErr) &&
			pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return fmt.Errorf(
				"%s-> on insert transfer: %w",
				op, erring.ErrAccountAlreadyExists,
			)
		}

		return fmt.Errorf(
			"%s-> on insert transfer: %w", op, err,
		)
	}

	err = r.accRepo.UpdateBalance(
		ctx,
		tx,
		trf.AccountOriginID,
		trf.AccountOrigin.Balance,
	)
	if err != nil {
		return fmt.Errorf(
			"%s-> on update account origin: %w",
			op,
			err,
		)
	}

	err = r.accRepo.UpdateBalance(
		ctx,
		tx,
		trf.AccountDestinationID,
		trf.AccountDestination.Balance,
	)
	if err != nil {
		return fmt.Errorf("%s-> on update account destination: %w", op, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s-> on commit transaction: %w", op, err)
	}

	return nil
}
