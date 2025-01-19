// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package accounts

import (
	"context"
	"sync"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			GetAllFunc: func(ctx context.Context) ([]entities.Account, error) {
//				panic("mock out the GetAll method")
//			},
//			GetByCPFFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
//				panic("mock out the GetByCPF method")
//			},
//			GetByIDFunc: func(ctx context.Context, id string) (entities.Account, error) {
//				panic("mock out the GetByID method")
//			},
//			InsertFunc: func(ctx context.Context, account entities.Account) error {
//				panic("mock out the Insert method")
//			},
//			UpdateBalanceFunc: func(ctx context.Context, accountID string, balance int) error {
//				panic("mock out the UpdateBalance method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) ([]entities.Account, error)

	// GetByCPFFunc mocks the GetByCPF method.
	GetByCPFFunc func(ctx context.Context, cpf string) (entities.Account, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, id string) (entities.Account, error)

	// InsertFunc mocks the Insert method.
	InsertFunc func(ctx context.Context, account entities.Account) error

	// UpdateBalanceFunc mocks the UpdateBalance method.
	UpdateBalanceFunc func(ctx context.Context, accountID string, balance int) error

	// calls tracks calls to the methods.
	calls struct {
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetByCPF holds details about calls to the GetByCPF method.
		GetByCPF []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Cpf is the cpf argument value.
			Cpf string
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// Insert holds details about calls to the Insert method.
		Insert []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Account is the account argument value.
			Account entities.Account
		}
		// UpdateBalance holds details about calls to the UpdateBalance method.
		UpdateBalance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AccountID is the accountID argument value.
			AccountID string
			// Balance is the balance argument value.
			Balance int
		}
	}
	lockGetAll        sync.RWMutex
	lockGetByCPF      sync.RWMutex
	lockGetByID       sync.RWMutex
	lockInsert        sync.RWMutex
	lockUpdateBalance sync.RWMutex
}

// GetAll calls GetAllFunc.
func (mock *RepositoryMock) GetAll(ctx context.Context) ([]entities.Account, error) {
	if mock.GetAllFunc == nil {
		panic("RepositoryMock.GetAllFunc: method is nil but Repository.GetAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedRepository.GetAllCalls())
func (mock *RepositoryMock) GetAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetByCPF calls GetByCPFFunc.
func (mock *RepositoryMock) GetByCPF(ctx context.Context, cpf string) (entities.Account, error) {
	if mock.GetByCPFFunc == nil {
		panic("RepositoryMock.GetByCPFFunc: method is nil but Repository.GetByCPF was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Cpf string
	}{
		Ctx: ctx,
		Cpf: cpf,
	}
	mock.lockGetByCPF.Lock()
	mock.calls.GetByCPF = append(mock.calls.GetByCPF, callInfo)
	mock.lockGetByCPF.Unlock()
	return mock.GetByCPFFunc(ctx, cpf)
}

// GetByCPFCalls gets all the calls that were made to GetByCPF.
// Check the length with:
//
//	len(mockedRepository.GetByCPFCalls())
func (mock *RepositoryMock) GetByCPFCalls() []struct {
	Ctx context.Context
	Cpf string
} {
	var calls []struct {
		Ctx context.Context
		Cpf string
	}
	mock.lockGetByCPF.RLock()
	calls = mock.calls.GetByCPF
	mock.lockGetByCPF.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *RepositoryMock) GetByID(ctx context.Context, id string) (entities.Account, error) {
	if mock.GetByIDFunc == nil {
		panic("RepositoryMock.GetByIDFunc: method is nil but Repository.GetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, id)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//
//	len(mockedRepository.GetByIDCalls())
func (mock *RepositoryMock) GetByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// Insert calls InsertFunc.
func (mock *RepositoryMock) Insert(ctx context.Context, account entities.Account) error {
	if mock.InsertFunc == nil {
		panic("RepositoryMock.InsertFunc: method is nil but Repository.Insert was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Account entities.Account
	}{
		Ctx:     ctx,
		Account: account,
	}
	mock.lockInsert.Lock()
	mock.calls.Insert = append(mock.calls.Insert, callInfo)
	mock.lockInsert.Unlock()
	return mock.InsertFunc(ctx, account)
}

// InsertCalls gets all the calls that were made to Insert.
// Check the length with:
//
//	len(mockedRepository.InsertCalls())
func (mock *RepositoryMock) InsertCalls() []struct {
	Ctx     context.Context
	Account entities.Account
} {
	var calls []struct {
		Ctx     context.Context
		Account entities.Account
	}
	mock.lockInsert.RLock()
	calls = mock.calls.Insert
	mock.lockInsert.RUnlock()
	return calls
}

// UpdateBalance calls UpdateBalanceFunc.
func (mock *RepositoryMock) UpdateBalance(ctx context.Context, accountID string, balance int) error {
	if mock.UpdateBalanceFunc == nil {
		panic("RepositoryMock.UpdateBalanceFunc: method is nil but Repository.UpdateBalance was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		AccountID string
		Balance   int
	}{
		Ctx:       ctx,
		AccountID: accountID,
		Balance:   balance,
	}
	mock.lockUpdateBalance.Lock()
	mock.calls.UpdateBalance = append(mock.calls.UpdateBalance, callInfo)
	mock.lockUpdateBalance.Unlock()
	return mock.UpdateBalanceFunc(ctx, accountID, balance)
}

// UpdateBalanceCalls gets all the calls that were made to UpdateBalance.
// Check the length with:
//
//	len(mockedRepository.UpdateBalanceCalls())
func (mock *RepositoryMock) UpdateBalanceCalls() []struct {
	Ctx       context.Context
	AccountID string
	Balance   int
} {
	var calls []struct {
		Ctx       context.Context
		AccountID string
		Balance   int
	}
	mock.lockUpdateBalance.RLock()
	calls = mock.calls.UpdateBalance
	mock.lockUpdateBalance.RUnlock()
	return calls
}
