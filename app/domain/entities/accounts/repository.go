package accounts

// Repository - Interface que define as assinaturas para o repository de accounts.
//go:generate moq -fmt goimports -out repository_mock.go . Repository
type Repository interface {
	ListAccounts() ([]Account, error)
	GetByID(id string) (Account, error)
	GetAccountByCPF(cpf string) (Account, error)
	Insert(account Account) error
	Update(account Account) error
}
