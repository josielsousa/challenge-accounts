package transfers

// Repository - Interface que define as assinaturas para o repository de transfers.
//go:generate moq -fmt goimports -out repository_mock.go . Repository
type Repository interface {
	Insert(transfer Transfer) error
	ListTransfers(accountID string) ([]Transfer, error)
}
