// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/types"
	"sync"
)

// Ensure, that accUsecaseMock does implement accUsecase.
// If this is not the case, regenerate this file with moq.
var _ accUsecase = &accUsecaseMock{}

// accUsecaseMock is a mock implementation of accUsecase.
//
//	func TestSomethingThatUsesaccUsecase(t *testing.T) {
//
//		// make and configure a mocked accUsecase
//		mockedaccUsecase := &accUsecaseMock{
//			CreateFunc: func(ctx context.Context, input accounts.AccountInput) (accounts.AccountOutput, error) {
//				panic("mock out the Create method")
//			},
//			GetAccountBalanceFunc: func(ctx context.Context, accountID string) (int, error) {
//				panic("mock out the GetAccountBalance method")
//			},
//			GetAllAccountsFunc: func(ctx context.Context) ([]accounts.AccountOutput, error) {
//				panic("mock out the GetAllAccounts method")
//			},
//		}
//
//		// use mockedaccUsecase in code that requires accUsecase
//		// and then make assertions.
//
//	}
type accUsecaseMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, input accounts.AccountInput) (accounts.AccountOutput, error)

	// GetAccountBalanceFunc mocks the GetAccountBalance method.
	GetAccountBalanceFunc func(ctx context.Context, accountID string) (int, error)

	// GetAllAccountsFunc mocks the GetAllAccounts method.
	GetAllAccountsFunc func(ctx context.Context) ([]accounts.AccountOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input accounts.AccountInput
		}
		// GetAccountBalance holds details about calls to the GetAccountBalance method.
		GetAccountBalance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AccountID is the accountID argument value.
			AccountID string
		}
		// GetAllAccounts holds details about calls to the GetAllAccounts method.
		GetAllAccounts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockCreate            sync.RWMutex
	lockGetAccountBalance sync.RWMutex
	lockGetAllAccounts    sync.RWMutex
}

// Create calls CreateFunc.
func (mock *accUsecaseMock) Create(ctx context.Context, input accounts.AccountInput) (accounts.AccountOutput, error) {
	if mock.CreateFunc == nil {
		panic("accUsecaseMock.CreateFunc: method is nil but accUsecase.Create was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Input accounts.AccountInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, input)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedaccUsecase.CreateCalls())
func (mock *accUsecaseMock) CreateCalls() []struct {
	Ctx   context.Context
	Input accounts.AccountInput
} {
	var calls []struct {
		Ctx   context.Context
		Input accounts.AccountInput
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// GetAccountBalance calls GetAccountBalanceFunc.
func (mock *accUsecaseMock) GetAccountBalance(ctx context.Context, accountID string) (int, error) {
	if mock.GetAccountBalanceFunc == nil {
		panic("accUsecaseMock.GetAccountBalanceFunc: method is nil but accUsecase.GetAccountBalance was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		AccountID string
	}{
		Ctx:       ctx,
		AccountID: accountID,
	}
	mock.lockGetAccountBalance.Lock()
	mock.calls.GetAccountBalance = append(mock.calls.GetAccountBalance, callInfo)
	mock.lockGetAccountBalance.Unlock()
	return mock.GetAccountBalanceFunc(ctx, accountID)
}

// GetAccountBalanceCalls gets all the calls that were made to GetAccountBalance.
// Check the length with:
//
//	len(mockedaccUsecase.GetAccountBalanceCalls())
func (mock *accUsecaseMock) GetAccountBalanceCalls() []struct {
	Ctx       context.Context
	AccountID string
} {
	var calls []struct {
		Ctx       context.Context
		AccountID string
	}
	mock.lockGetAccountBalance.RLock()
	calls = mock.calls.GetAccountBalance
	mock.lockGetAccountBalance.RUnlock()
	return calls
}

// GetAllAccounts calls GetAllAccountsFunc.
func (mock *accUsecaseMock) GetAllAccounts(ctx context.Context) ([]accounts.AccountOutput, error) {
	if mock.GetAllAccountsFunc == nil {
		panic("accUsecaseMock.GetAllAccountsFunc: method is nil but accUsecase.GetAllAccounts was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAllAccounts.Lock()
	mock.calls.GetAllAccounts = append(mock.calls.GetAllAccounts, callInfo)
	mock.lockGetAllAccounts.Unlock()
	return mock.GetAllAccountsFunc(ctx)
}

// GetAllAccountsCalls gets all the calls that were made to GetAllAccounts.
// Check the length with:
//
//	len(mockedaccUsecase.GetAllAccountsCalls())
func (mock *accUsecaseMock) GetAllAccountsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAllAccounts.RLock()
	calls = mock.calls.GetAllAccounts
	mock.lockGetAllAccounts.RUnlock()
	return calls
}

// Ensure, that authUsecaseMock does implement authUsecase.
// If this is not the case, regenerate this file with moq.
var _ authUsecase = &authUsecaseMock{}

// authUsecaseMock is a mock implementation of authUsecase.
//
//	func TestSomethingThatUsesauthUsecase(t *testing.T) {
//
//		// make and configure a mocked authUsecase
//		mockedauthUsecase := &authUsecaseMock{
//			SigninFunc: func(ctx context.Context, credential types.Credentials) (types.Auth, error) {
//				panic("mock out the Signin method")
//			},
//		}
//
//		// use mockedauthUsecase in code that requires authUsecase
//		// and then make assertions.
//
//	}
type authUsecaseMock struct {
	// SigninFunc mocks the Signin method.
	SigninFunc func(ctx context.Context, credential types.Credentials) (types.Auth, error)

	// calls tracks calls to the methods.
	calls struct {
		// Signin holds details about calls to the Signin method.
		Signin []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Credential is the credential argument value.
			Credential types.Credentials
		}
	}
	lockSignin sync.RWMutex
}

// Signin calls SigninFunc.
func (mock *authUsecaseMock) Signin(ctx context.Context, credential types.Credentials) (types.Auth, error) {
	if mock.SigninFunc == nil {
		panic("authUsecaseMock.SigninFunc: method is nil but authUsecase.Signin was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		Credential types.Credentials
	}{
		Ctx:        ctx,
		Credential: credential,
	}
	mock.lockSignin.Lock()
	mock.calls.Signin = append(mock.calls.Signin, callInfo)
	mock.lockSignin.Unlock()
	return mock.SigninFunc(ctx, credential)
}

// SigninCalls gets all the calls that were made to Signin.
// Check the length with:
//
//	len(mockedauthUsecase.SigninCalls())
func (mock *authUsecaseMock) SigninCalls() []struct {
	Ctx        context.Context
	Credential types.Credentials
} {
	var calls []struct {
		Ctx        context.Context
		Credential types.Credentials
	}
	mock.lockSignin.RLock()
	calls = mock.calls.Signin
	mock.lockSignin.RUnlock()
	return calls
}

// Ensure, that trfUsecaseMock does implement trfUsecase.
// If this is not the case, regenerate this file with moq.
var _ trfUsecase = &trfUsecaseMock{}

// trfUsecaseMock is a mock implementation of trfUsecase.
//
//	func TestSomethingThatUsestrfUsecase(t *testing.T) {
//
//		// make and configure a mocked trfUsecase
//		mockedtrfUsecase := &trfUsecaseMock{
//			DoTransferFunc: func(ctx context.Context, input transfers.TransferInput) (transfers.TransferOutput, error) {
//				panic("mock out the DoTransfer method")
//			},
//			ListTransfersAccountFunc: func(ctx context.Context, accOriginID string) ([]transfers.TransferOutput, error) {
//				panic("mock out the ListTransfersAccount method")
//			},
//		}
//
//		// use mockedtrfUsecase in code that requires trfUsecase
//		// and then make assertions.
//
//	}
type trfUsecaseMock struct {
	// DoTransferFunc mocks the DoTransfer method.
	DoTransferFunc func(ctx context.Context, input transfers.TransferInput) (transfers.TransferOutput, error)

	// ListTransfersAccountFunc mocks the ListTransfersAccount method.
	ListTransfersAccountFunc func(ctx context.Context, accOriginID string) ([]transfers.TransferOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// DoTransfer holds details about calls to the DoTransfer method.
		DoTransfer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input transfers.TransferInput
		}
		// ListTransfersAccount holds details about calls to the ListTransfersAccount method.
		ListTransfersAccount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AccOriginID is the accOriginID argument value.
			AccOriginID string
		}
	}
	lockDoTransfer           sync.RWMutex
	lockListTransfersAccount sync.RWMutex
}

// DoTransfer calls DoTransferFunc.
func (mock *trfUsecaseMock) DoTransfer(ctx context.Context, input transfers.TransferInput) (transfers.TransferOutput, error) {
	if mock.DoTransferFunc == nil {
		panic("trfUsecaseMock.DoTransferFunc: method is nil but trfUsecase.DoTransfer was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Input transfers.TransferInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockDoTransfer.Lock()
	mock.calls.DoTransfer = append(mock.calls.DoTransfer, callInfo)
	mock.lockDoTransfer.Unlock()
	return mock.DoTransferFunc(ctx, input)
}

// DoTransferCalls gets all the calls that were made to DoTransfer.
// Check the length with:
//
//	len(mockedtrfUsecase.DoTransferCalls())
func (mock *trfUsecaseMock) DoTransferCalls() []struct {
	Ctx   context.Context
	Input transfers.TransferInput
} {
	var calls []struct {
		Ctx   context.Context
		Input transfers.TransferInput
	}
	mock.lockDoTransfer.RLock()
	calls = mock.calls.DoTransfer
	mock.lockDoTransfer.RUnlock()
	return calls
}

// ListTransfersAccount calls ListTransfersAccountFunc.
func (mock *trfUsecaseMock) ListTransfersAccount(ctx context.Context, accOriginID string) ([]transfers.TransferOutput, error) {
	if mock.ListTransfersAccountFunc == nil {
		panic("trfUsecaseMock.ListTransfersAccountFunc: method is nil but trfUsecase.ListTransfersAccount was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		AccOriginID string
	}{
		Ctx:         ctx,
		AccOriginID: accOriginID,
	}
	mock.lockListTransfersAccount.Lock()
	mock.calls.ListTransfersAccount = append(mock.calls.ListTransfersAccount, callInfo)
	mock.lockListTransfersAccount.Unlock()
	return mock.ListTransfersAccountFunc(ctx, accOriginID)
}

// ListTransfersAccountCalls gets all the calls that were made to ListTransfersAccount.
// Check the length with:
//
//	len(mockedtrfUsecase.ListTransfersAccountCalls())
func (mock *trfUsecaseMock) ListTransfersAccountCalls() []struct {
	Ctx         context.Context
	AccOriginID string
} {
	var calls []struct {
		Ctx         context.Context
		AccOriginID string
	}
	mock.lockListTransfersAccount.RLock()
	calls = mock.calls.ListTransfersAccount
	mock.lockListTransfersAccount.RUnlock()
	return calls
}
