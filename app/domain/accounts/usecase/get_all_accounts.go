package usecase

import (
	"context"
	"fmt"

	accUC "github.com/josielsousa/challenge-accounts/app/domain/accounts"
)

func (a Account) GetAllAccounts(ctx context.Context) ([]accUC.AccountOutput, error) {
	const op = `accounts.GetAllAccounts`

	accs, err := a.accRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s -> %s: %w", op, "on get all accounts", err)
	}

	out := make([]accUC.AccountOutput, 0, len(accs))
	for _, acc := range accs {
		out = append(out, accUC.AccountOutput{
			ID:        acc.ID,
			Name:      acc.Name,
			Balance:   acc.Balance,
			CPF:       acc.CPF,
			CreatedAt: acc.CreatedAt,
			UpdatedAt: acc.UpdatedAt,
		})
	}

	return out, nil
}
