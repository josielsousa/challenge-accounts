package transfers

import (
	"context"
	"time"
)

type Usecase interface {
	DoTransfer(ctx context.Context, input TransferInput) error
	ListTransfersAccount(ctx context.Context, accOriginID string) ([]TransferOutput, error)
}

type TransferInput struct {
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
}

type TransferOutput struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
}
