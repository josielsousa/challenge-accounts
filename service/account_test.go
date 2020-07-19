package service

import (
	"testing"
	"time"

	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/repo/model/mocks"
	"github.com/josielsousa/challenge-accounts/types"
)

var (
	stg         *mocks.MockAccountStorage
	logAcc      *types.MockAPILogProvider
	service     *AccountService
	accountTest model.Account
)

func mockLog() {
	logAcc = &types.MockAPILogProvider{}
	logAcc.OnInfo = func(info string) {}
	logAcc.OnError = func(trace string, erro error) {}
}

func mockAccountStorage() model.AccountStorage {
	accountTest = model.Account{
		ID:        "000-0000000-000",
		Cpf:       "123.456.789",
		Secret:    "xxxx",
		Name:      "Teste",
		Ballance:  3.99,
		CreatedAt: time.Now(),
	}

	stg = &mocks.MockAccountStorage{
		OnGetAccount: func(id string) (*model.Account, error) {
			return &accountTest, nil
		},

		OnUpdate: func(account model.Account) (*model.Account, error) {
			return &accountTest, nil
		},

		OnInsert: func(account model.Account) (*model.Account, error) {
			return &accountTest, nil
		},

		OnGetAllAccounts: func() (accounts []model.Account, err error) {
			accounts = make([]model.Account, 0)
			accounts = append(accounts, accountTest)
			return accounts, nil
		},
	}

	return stg
}

func setup() *AccountService {
	mockLog()
	mockAccountStorage()
	return NewAccountService(stg, logAcc)
}

func TestService_InsertAccount(t *testing.T) {
	service = setup()

	t.Run("Teste Inserir account sucesso", func(t *testing.T) {

	})
}
