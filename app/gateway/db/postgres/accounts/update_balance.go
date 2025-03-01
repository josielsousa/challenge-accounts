package accounts

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

func (*Repository) UpdateBalance(
	ctx context.Context,
	tx pgx.Tx,
	accountID string,
	balance int,
) error {
	const operation = `Repository.Accounts.UpdateBalance`

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
		return fmt.Errorf(
			"%s-> on update balance account: %w", operation, err,
		)
	}

	if cmTag.RowsAffected() != 1 {
		return fmt.Errorf(
			"%s-> on check rows affected: %w",
			operation, erring.ErrUpdateAccountNotPerformed,
		)
	}

	return nil
}
