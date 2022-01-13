package accounts

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

const (
	queryByCPF = `
		SELECT 
			id,
			name,
			cpf,
			secret,
			balance,
			created_at,
			updated_at
		FROM accounts 
		WHERE cpf = $1 
	`

	queryByID = `
		SELECT 
			id,
			name,
			cpf,
			secret,
			balance,
			created_at,
			updated_at
		FROM accounts 
		WHERE id = $1 
	`
)

func (r *Repository) GetByCPF(ctx context.Context, numCPF string) (accounts.Account, error) {
	const op = `Repository.Accounts.GetAccountByCPF`

	acc, err := r.getAccount(ctx, numCPF, queryByCPF)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("%s-> %s: %w", op, "on query by CPF", err)
	}

	return acc, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (accounts.Account, error) {
	const op = `Repository.Accounts.GetAccountByCPF`

	acc, err := r.getAccount(ctx, id, queryByID)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("%s-> %s: %w", op, "on query by id", err)
	}

	return acc, nil
}

func (r *Repository) getAccount(ctx context.Context, param string, query string) (accounts.Account, error) {
	const op = `Repository.Accounts.getAccount`

	row := r.db.QueryRow(
		ctx,
		query,
		param,
	)

	var acc accounts.Account

	err := row.Scan(
		&acc.ID,
		&acc.Name,
		&acc.CPF,
		&acc.Secret,
		&acc.Balance,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("%s-> %s: %w", op, "on get account by cpf", err)
	}

	return acc, nil
}
