package accounts

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

const (
	queryGetAllAccs = `
		SELECT
			id,
			name,
			cpf,
			secret,
			balance,
			created_at,
			updated_at
		FROM accounts
	`
)

func (r *Repository) GetAll(ctx context.Context) ([]accounts.Account, error) {
	const operation = `Repository.Accounts.GetAll`

	rows, err := r.db.Query(
		ctx,
		queryGetAllAccs,
	)
	if err != nil {
		return nil, fmt.Errorf("%s -> %s: %w", operation, "on get all accounts", err)
	}

	defer rows.Close()

	accs := make([]accounts.Account, 0)

	for rows.Next() {
		var acc accounts.Account

		err := rows.Scan(
			&acc.ID,
			&acc.Name,
			&acc.CPF,
			&acc.Secret,
			&acc.Balance,
			&acc.CreatedAt,
			&acc.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s-> %s: %w", operation, "on scan get all accounts", err)
		}

		accs = append(accs, acc)
	}

	return accs, nil
}
