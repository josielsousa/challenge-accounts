package usecase

import (
	"context"
	"fmt"
)

func (a Account) GetAccountBalance(ctx context.Context, accountID string) (int, error) {
	const op = `Usecase.Account.GetAccountBalance`

	acc, err := a.accRepo.GetByID(ctx, accountID)
	if err != nil {
		return 0, fmt.Errorf("%s -> %s: %w", op, "on get account by id", err)
	}

	return acc.Balance, nil
}
