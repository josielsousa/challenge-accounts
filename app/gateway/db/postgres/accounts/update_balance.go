package accounts

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

func (*Repository) UpdateBalance(ctx context.Context, tx pgx.Tx, accountID string, balance int) error {
	const op = `Repository.Accounts.UpdateBalance`

	query := `
        UPDATE accounts SET
			balance = $2,
			updated_at = now()
		WHERE id = $1
    `

	cmTag, err := tx.Exec(
		ctx,
		query,
		accountID,
		balance,
	)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on update balance account", err)
	}

	if cmTag.RowsAffected() != 1 {
		return fmt.Errorf("%s-> %s: %w", op, "on check rows affected", accounts.ErrUpdateAccountNotPerformed)
	}

	return nil
}
