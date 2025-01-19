package accounts

import (
	"context"
	"fmt"
)

func (a Account) GetAccountBalance(ctx context.Context, accountID string) (int, error) {
	const op = `accounts.GetAccountBalance`

	acc, err := a.R.GetByID(ctx, accountID)
	if err != nil {
		return 0, fmt.Errorf("%s -> %s: %w", op, "on get account by id", err)
	}

	return acc.Balance, nil
}
