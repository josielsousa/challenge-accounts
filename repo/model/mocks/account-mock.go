package mocks

import "github.com/josielsousa/challenge-accounts/repo/model"

// MockAccountStorage - Mock account storage
type MockAccountStorage struct {
	OnGetAllAccounts func() ([]string, error)
	OnGetAccount     func(id string) (*model.Account, error)
	OnInsert         func(account model.Account) (*model.Account, error)
	OnUpdate         func(account model.Account) (*model.Account, error)
}

// GetAllAccounts - Mock provider recuperar todas as accounts.
func (m *MockAccountStorage) GetAllAccounts() ([]string, error) {
	return m.OnGetAllAccounts()
}

// GetAccount - Mock provider recuperar uma account por id.
func (m *MockAccountStorage) GetAccount(id string) (*model.Account, error) {
	return m.OnGetAccount(id)
}

// Insert - Mock provider para criar uma account.
func (m *MockAccountStorage) Insert(account model.Account) (*model.Account, error) {
	return m.OnInsert(account)
}

// Update - Mock provider para atualizar uma account.
func (m *MockAccountStorage) Update(account model.Account) (*model.Account, error) {
	return m.OnUpdate(account)
}
