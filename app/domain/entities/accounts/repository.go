package accounts

import "context"

// Repository - Interface que define as assinaturas para o repository de accounts.
//go:generate moq -fmt goimports -out repository_mock.go . Repository
type Repository interface {
	GetAll(ctx context.Context) ([]Account, error)
	GetByID(ctx context.Context, id string) (Account, error)
	GetByCPF(ctx context.Context, cpf string) (Account, error)
	Insert(ctx context.Context, account Account) error
	Update(ctx context.Context, account Account) error
}
