package dbgorm_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/josielsousa/challenge-accounts/repo/dbgorm"
	"github.com/josielsousa/challenge-accounts/types"
	_ "github.com/mattn/go-sqlite3"

	"github.com/josielsousa/challenge-accounts/repo/model"
)

//Constante de mensagens
const (
	ErrorInsertAccount            = "Error on insert account"
	ErrorUpdateAccount            = "Error on update account"
	ErrorGetAccount               = "Error on get an account by id"
	ErrorGetAllAccount            = "Error on get all accounts"
	ErrorInsertAccountEmptyReturn = "Error on insert account, empty return"
	ErrorInsertAccountIdDiffer    = "Error on insert account, id differ returning"
	ErrorUpdateAccountNameDiffer  = "Error on update, name returns differ original name"
	ErrorGetAccountIdDiffer       = "Error on get account, id returns differ original id"
	ErrorGetAccountsEmptyReturn   = "Error on get all accounts, no content"
)

var (
	dbAcc       *gorm.DB
	accountTest model.Account
	stgAcc      *dbgorm.AccountStorage
)

func setupTestAccounts(t *testing.T) *dbgorm.AccountStorage {
	if dbAcc == nil {
		dbGorm, err := gorm.Open("sqlite3", types.DatabaseNameTest)
		if err != nil {
			t.Error(types.ErrorOpenConnection, err)
			return nil
		}

		dbAcc = dbGorm
	}

	accountTest = model.Account{
		ID:   uuid.New().String(),
		Name: "Teste Account",
	}

	return dbgorm.NewAccountStorage(dbAcc)
}

func TestStorageInsertAccount(t *testing.T) {
	stgAcc = setupTestAccounts(t)

	t.Run("Teste Inserir account sucesso", func(t *testing.T) {
		acc, err := stgAcc.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountIdDiffer)
		}
	})
}

func TestStorageUpdateAccount(t *testing.T) {
	stgAcc = setupTestAccounts(t)

	t.Run("Teste Update account sucesso", func(t *testing.T) {
		acc, err := stgAcc.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountEmptyReturn, err)
		}

		newName := "Novo nome"
		acc.Name = newName

		accUpdate, err := stgAcc.Update(*acc)
		if err != nil {
			t.Error(ErrorUpdateAccount, err)
		}

		if acc.Name != accUpdate.Name {
			t.Error(ErrorUpdateAccountNameDiffer, err)
		}
	})
}

func TestStorageGetAccount(t *testing.T) {
	stgAcc = setupTestAccounts(t)

	t.Run("Teste Get Account sucesso", func(t *testing.T) {
		acc, err := stgAcc.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountEmptyReturn, err)
		}

		accByID, err := stgAcc.GetAccount(acc.ID)
		if err != nil {
			t.Error(ErrorGetAccount, err)
		}

		if acc.ID != accByID.ID {
			t.Error(ErrorUpdateAccountNameDiffer, err)
		}
	})
}

func TestStorageGetAllAccounts(t *testing.T) {
	stgAcc = setupTestAccounts(t)

	t.Run("Teste Get Account sucesso", func(t *testing.T) {
		acc, err := stgAcc.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountEmptyReturn, err)
		}

		allAccounts, err := stgAcc.GetAllAccounts()
		if err != nil {
			t.Error(ErrorGetAllAccount, err)
		}

		if len(allAccounts) <= 0 {
			t.Error(ErrorGetAccountsEmptyReturn)
		}
	})
}
