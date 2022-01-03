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
)

func (r *Repository) Insert(ctx context.Context, acc accounts.Account) error {
	const op = `Repository.Accounts.Insert`
	if len(acc.ID) <= 0 {
		acc.ID = uuid.NewString()
	}

	acc.CreatedAt = time.Now().In(time.Local)

	query := `
        INSERT INTO accounts(
            id,
            name,
            cpf,
            secret,
            balance,
            created_at,
        ) VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at
    `

	err := r.db.QueryRow(
		ctx,
		query,
		acc.ID,
		acc.Name,
		acc.CPF,
		acc.Secret,
		acc.Balance,
		acc.CreatedAt,
	).Scan(
		&acc.ID,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return fmt.Errorf("%s-> %s:%w", op, "on insert account", accounts.ErrAccountAlreadyExists)
		}

		return fmt.Errorf("%s-> %s:%w", op, "on insert account", err)
	}

	return nil
}
