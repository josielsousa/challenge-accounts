package accounts

import (
	"context"
	"fmt"
)

func (u Usecase) GetAccountBalance(
	ctx context.Context, accountID string,
) (int, error) {
	const op = `accounts.GetAccountBalance`

	acc, err := u.R.GetByID(ctx, accountID)
	if err != nil {
		return 0, fmt.Errorf("%s -> on get account by id: %w", op, err)
	}

	return acc.Balance, nil
}
