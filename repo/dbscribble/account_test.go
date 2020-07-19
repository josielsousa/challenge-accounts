package dbscribble_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/josielsousa/challenge-accounts/repo/dbscribble"
	"github.com/josielsousa/challenge-accounts/repo/model"
	scribble "github.com/nanobox-io/golang-scribble"
)

//Constante de mensagens
const (
	DatabaseNameTest              = "../../DBTestAccount"
	ErrorInsertAccount            = "Error on insert account"
	ErrorUpdateAccount            = "Error on update account"
	ErrorGetAccount               = "Error on get an account by id"
	ErrorGetAllAccount            = "Error on get all accounts"
	ErrorInsertAccountEmptyReturn = "Error on insert account, empty return"
	ErrorUpdateAccountNameDiffer  = "Error on update, name returns differ original name"
	ErrorGetAccountIdDiffer       = "Error on get account, id returns differ original id"
	ErrorGetAccountsEmptyReturn   = "Error on get all accounts, no content"
)

var (
	stg         *dbscribble.AccountStorage
	db          *scribble.Driver
	accountTest model.Account
)

func setup() *dbscribble.AccountStorage {
	if db == nil {
		dbScribble, err := scribble.New(DatabaseNameTest, nil)
		if err != nil {
			return nil
		}

		db = dbScribble
	}

	accountTest = model.Account{
		ID:   uuid.New().String(),
		Name: "Teste Account",
	}

	return dbscribble.NewAccountStorage(db)
}

func TestStorageInsertAccount(t *testing.T) {
	stg = setup()

	t.Run("Teste Inserir account sucesso", func(t *testing.T) {
		acc, err := stg.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountEmptyReturn, err)
		}
	})
}

func TestStorageUpdateAccount(t *testing.T) {
	stg = setup()

	t.Run("Teste Update account sucesso", func(t *testing.T) {
		acc, err := stg.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountEmptyReturn, err)
		}

		newName := "Novo nome"
		acc.Name = newName

		accUpdate, err := stg.Update(*acc)
		if err != nil {
			t.Error(ErrorUpdateAccount, err)
		}

		if acc.Name != accUpdate.Name {
			t.Error(ErrorUpdateAccountNameDiffer, err)
		}
	})
}

func TestStorageGetAccount(t *testing.T) {
	stg = setup()

	t.Run("Teste Get Account sucesso", func(t *testing.T) {
		acc, err := stg.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountEmptyReturn, err)
		}

		accByID, err := stg.GetAccount(acc.ID)
		if err != nil {
			t.Error(ErrorGetAccount, err)
		}

		if acc.ID != accByID.ID {
			t.Error(ErrorUpdateAccountNameDiffer, err)
		}
	})
}

func TestStorageGetAllAccounts(t *testing.T) {
	stg = setup()

	t.Run("Teste Get Account sucesso", func(t *testing.T) {
		acc, err := stg.Insert(accountTest)
		if err != nil {
			t.Error(ErrorInsertAccount, err)
		}

		if acc.ID != accountTest.ID {
			t.Error(ErrorInsertAccountEmptyReturn, err)
		}

		allAccounts, err := stg.GetAllAccounts()
		if err != nil {
			t.Error(ErrorGetAllAccount, err)
		}

		if len(allAccounts) <= 0 {
			t.Error(ErrorGetAccountsEmptyReturn)
		}
	})
}
