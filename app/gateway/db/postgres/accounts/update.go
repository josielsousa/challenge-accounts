package accounts

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

func (*Repository) Update(ctx context.Context, tx pgx.Tx, acc accounts.Account) error {
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
    `

	cmTag, err := tx.Exec(
		ctx,
		query,
		acc.ID,
		acc.Name,
		acc.CPF.Value(),
		acc.Secret.Value(),
		acc.Balance,
		acc.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on update account", err)
	}

	if cmTag.RowsAffected() != 1 {
		return fmt.Errorf("%s-> %s: %w", op, "on check rows affected", accounts.ErrUpdateAccountNotPerformed)
	}

	return nil
}
