package service

import (
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
)

// AccountService - Implementação do service para as accounts.
type AccountService struct {
	stgAccount model.AccountStorage
	logger     types.APILogProvider
}

//NewAccountService - Instância o service com a dependência `log` inicializada.
func NewAccountService(stgAccount model.AccountStorage, log types.APILogProvider) *AccountService {
	return &AccountService{
		logger:     log,
		stgAccount: stgAccount,
	}
}
