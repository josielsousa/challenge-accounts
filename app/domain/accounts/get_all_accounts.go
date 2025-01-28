package accounts

import (
	"context"
	"fmt"
)

func (u Usecase) GetAllAccounts(ctx context.Context) ([]AccountOutput, error) {
	const op = `accounts.GetAllAccounts`

	accs, err := u.R.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s -> on get all accounts: %w", op, err)
	}

	out := make([]AccountOutput, 0, len(accs))

	for _, acc := range accs {
		out = append(out, AccountOutput{
			ID:        acc.ID,
			Name:      acc.Name,
			Balance:   acc.Balance,
			CPF:       acc.CPF.Value(),
			CreatedAt: acc.CreatedAt,
			UpdatedAt: acc.UpdatedAt,
		})
	}

	return out, nil
}
