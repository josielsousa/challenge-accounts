package accounts

import (
	"context"
	"fmt"
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

func (r *Repository) Update(ctx context.Context, acc accounts.Account) error {
	const op = `Repository.Accounts.Update`

	if acc.UpdatedAt.IsZero() {
		acc.UpdatedAt = time.Now().In(time.Local)
	}

	query := `
        UPDATE accounts SET
			name = $2,
			cpf = $3,
			secret = $4,
			balance = $5,
			updated_at = $6
		WHERE id = $1
        RETURNING id, created_at, updated_at
    `

	row := r.db.QueryRow(
		ctx,
		query,
		acc.ID,
		acc.Name,
		acc.CPF.String(),
		acc.Secret.String(),
		acc.Balance,
		acc.UpdatedAt,
	)

	err := row.Scan(
		&acc.ID,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s-> %s:%w", op, "on update account", err)
	}

	return nil
}
