package accounts

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

func (r *Repository) Insert(ctx context.Context, acc entities.Account) error {
	const op = `Repository.Accounts.Insert`

	if len(acc.ID) == 0 {
		acc.ID = uuid.NewString()
	}

	if acc.CreatedAt.IsZero() {
		acc.CreatedAt = time.Now().In(time.UTC)
	}

	query := `
        INSERT INTO accounts(
            id,
            name,
            cpf,
            secret,
            balance,
            created_at
        ) VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at
    `

	row := r.db.QueryRow(
		ctx,
		query,
		acc.ID,
		acc.Name,
		acc.CPF.Value(),
		acc.Secret.Value(),
		acc.Balance,
		acc.CreatedAt,
	)

	err := row.Scan(
		&acc.ID,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return fmt.Errorf(
				"%s-> on insert account: %w",
				op, erring.ErrAccountAlreadyExists,
			)
		}

		return fmt.Errorf("%s-> on insert account: %w", op, err)
	}

	return nil
}
