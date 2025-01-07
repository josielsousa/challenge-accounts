package mocks

import "github.com/josielsousa/challenge-accounts/repo/model"

// MockTransferStorage - Mock transfer storage.
type MockTransferStorage struct {
	OnGetAllTransfers func(accountID string) ([]model.Transfer, error)
	OnInsert          func(transfer model.Transfer) (*model.Transfer, error)
}

// GetAllTransfers - Mock provider recuperar todas as transfers.
func (m *MockTransferStorage) GetAllTransfers(accountID string) ([]model.Transfer, error) {
	return m.OnGetAllTransfers(accountID)
}

// Insert - Mock provider para criar uma transfer.
func (m *MockTransferStorage) Insert(transfer model.Transfer) (*model.Transfer, error) {
	return m.OnInsert(transfer)
}
