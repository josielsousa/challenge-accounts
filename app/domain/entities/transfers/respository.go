package transfers

import "context"

// Repository - Interface que define as assinaturas para o repository de transfers.
//go:generate moq -fmt goimports -out repository_mock.go . Repository
type Repository interface {
	Insert(ctx context.Context, transfer TransferData) error
	ListTransfers(ctx context.Context, accOriginID string) ([]Transfer, error)
}
