package accounts

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
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
	const op = `Repository.Accounts.GetAll`

	rows, err := r.db.Query(
		ctx,
		queryGetAllAccs,
	)
	if err != nil {
		return nil, fmt.Errorf("%s -> %s: %w", op, "on get all accounts", err)
	}

	defer rows.Close()

	accs := make([]accounts.Account, 0)

	for rows.Next() {
		var (
			numCPF string
			sec    string
			acc    accounts.Account
		)

		err := rows.Scan(
			&acc.ID,
			&acc.Name,
			&numCPF,
			&sec,
			&acc.Balance,
			&acc.CreatedAt,
			&acc.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s-> %s: %w", op, "on scan get all accounts", err)
		}

		if len(numCPF) > 0 {
			accCPF, err := cpf.NewCPF(numCPF)
			if err != nil {
				return nil, fmt.Errorf("%s-> %s: %w", op, "on new CPF vos", err)
			}

			acc.CPF = accCPF
		}

		acc.SetHashedSecret(sec)
		accs = append(accs, acc)
	}

	return accs, nil
}
