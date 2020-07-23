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
	ErrorInsertTransfer            = "Error on insert transfer"
	ErrorUpdateTransfer            = "Error on update transfer"
	ErrorGetTransfer               = "Error on get an transfer by id"
	ErrorGetAllTransfer            = "Error on get all transfers"
	ErrorInsertTransferEmptyReturn = "Error on insert transfer, empty return"
	ErrorInsertTransferIdDiffer    = "Error on insert transfer, id differ returning"
	ErrorUpdateTransferNameDiffer  = "Error on update, name returns differ original name"
	ErrorGetTransferIdDiffer       = "Error on get transfer, id returns differ original id"
	ErrorGetTransfersEmptyReturn   = "Error on get all transfers, no content"
)

var (
	dbTransfer   *gorm.DB
	transferTest model.Transfer
	stgTransfer  *dbgorm.TransferStorage
)

func setupTestTransfers(t *testing.T) *dbgorm.TransferStorage {
	if dbTransfer == nil {
		dbGorm, err := gorm.Open("sqlite3", types.DatabaseNameTest)
		if err != nil {
			t.Error(types.ErrorOpenConnection, err)
			return nil
		}

		dbTransfer = dbGorm
	}

	transferTest = model.Transfer{
		ID:                   "",
		Amount:               3.50,
		AccountOriginID:      uuid.New().String(),
		AccountDestinationID: uuid.New().String(),
	}

	return dbgorm.NewTransferStorage(dbTransfer)
}

func TestStorageInsertTransfer(t *testing.T) {
	stgTransfer = setupTestTransfers(t)

	t.Run("Teste Inserir transfer sucesso", func(t *testing.T) {
		trf, err := stgTransfer.Insert(transferTest)
		if err != nil {
			t.Error(ErrorInsertTransfer, err)
		}

		if len(trf.ID) <= 0 {
			t.Error(ErrorInsertTransferIdDiffer)
		}
	})
}

func TestStorageGetAllTransfers(t *testing.T) {
	stgTransfer = setupTestTransfers(t)

	t.Run("Teste Get Transfer sucesso", func(t *testing.T) {
		transfer, err := stgTransfer.Insert(transferTest)
		if err != nil {
			t.Error(ErrorInsertTransfer, err)
		}

		if len(transfer.ID) <= 0 {
			t.Error(ErrorInsertTransferEmptyReturn, err)
		}

		allTransfers, err := stgTransfer.GetAllTransfers()
		if err != nil {
			t.Error(ErrorGetAllTransfer, err)
		}

		if len(allTransfers) <= 0 {
			t.Error(ErrorGetTransfersEmptyReturn)
		}
	})
}
