package usecase

import (
	"context"

	trfUC "github.com/josielsousa/challenge-accounts/app/domain/transfers"
)

func (t Transfer) ListTransfersAccount(ctx context.Context, accOriginID string) ([]trfUC.TransferOutput, error) {
	return nil, nil
}
