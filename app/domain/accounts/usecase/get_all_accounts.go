package usecase

import (
	"context"

	accUC "github.com/josielsousa/challenge-accounts/app/domain/accounts"
)

func (a Account) GetAllAccounts(ctx context.Context) ([]accUC.AccountOutput, error) {
	return nil, nil
}
