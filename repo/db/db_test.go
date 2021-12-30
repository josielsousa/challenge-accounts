package db_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/josielsousa/challenge-accounts/repo/db"
)

// Constante de mensagens
const (
	DatabaseType                             = "nope"
	ErrorOpenTransaction                     = "Error on open transaction"
	ErrorOpenConnection                      = "Error on open connection"
	ErrorOpenConnectionNotImplemented        = "Error on open connection not implemented"
	ErrorOpenConnectionServiceNil            = "Error on open connection, service nil"
	ErrorOpenConnectionNotImplementedService = "Error on open connection not implemented, service is not nil"
)

func TestDbOpenGormSuccess(t *testing.T) {
	t.Run("Teste Abrir conexão com o banco de dados utilizando Gorm", func(t *testing.T) {
		srv, err := db.Open(db.Gorm)
		if err != nil {
			t.Error(ErrorOpenConnection, err)
		}

		if srv == nil {
			t.Error(ErrorOpenConnectionServiceNil)
		}
	})
}

func TestDbOpenFail(t *testing.T) {
	t.Run("Teste Abrir conexão com o banco de dados, fail", func(t *testing.T) {
		srv, err := db.Open(DatabaseType)
		if err == nil {
			t.Error(ErrorOpenConnectionNotImplemented)
		}

		if srv != nil {
			t.Error(ErrorOpenConnectionNotImplementedService)
		}
	})
}

func TestDbCloseGormSuccess(t *testing.T) {
	t.Run("Teste Abrir conexão com o banco de dados utilizando Gorm", func(t *testing.T) {
		srv, err := db.Open(db.Gorm)
		if err != nil {
			t.Error(ErrorOpenConnection, err)
		}

		srv.Close()
	})
}

func TestDbRollbackGormSuccess(t *testing.T) {
	t.Run("Teste Abrir conexão com o banco de dados utilizando Gorm", func(t *testing.T) {
		srv, err := db.Open(db.Gorm)
		if err != nil {
			t.Error(ErrorOpenConnection, err)
		}

		srv.Rollback()
		srv.Close()
	})
}

func TestDbCommitGormSuccess(t *testing.T) {
	t.Run("Teste Abrir conexão com o banco de dados utilizando Gorm", func(t *testing.T) {
		srv, err := db.Open(db.Gorm)
		if err != nil {
			t.Error(ErrorOpenConnection, err)
		}

		srv.Commit()
		srv.Close()
	})
}

func TestDbOpenTransactionGormSuccess(t *testing.T) {
	t.Run("Teste Abrir conexão com o banco de dados utilizando Gorm", func(t *testing.T) {
		srv, err := db.Open(db.Gorm)
		if err != nil {
			t.Error(ErrorOpenConnection, err)
		}

		tx := srv.BeginTransaction()
		if tx == nil {
			t.Error(ErrorOpenTransaction, err)
		}

		tx.Rollback()
		srv.Commit()
		srv.Close()
	})
}
