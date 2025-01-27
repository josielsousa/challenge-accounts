package accounts

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
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

func (r *Repository) GetAll(ctx context.Context) ([]entities.Account, error) {
	const operation = `Repository.Accounts.GetAll`

	rows, err := r.db.Query(
		ctx,
		queryGetAllAccs,
	)
	if err != nil {
		return nil, fmt.Errorf("%s -> on get all accounts: %w", operation, err)
	}

	defer rows.Close()

	accs := make([]entities.Account, 0)

	for rows.Next() {
		var acc entities.Account

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
			return nil, fmt.Errorf(
				"%s-> on scan get all accounts: %w", operation, err,
			)
		}

		accs = append(accs, acc)
	}

	return accs, nil
}
