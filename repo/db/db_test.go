package db_test

import (
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/josielsousa/challenge-accounts/repo/db"
)

//Constante de mensagens
const (
	DatabaseType                             = "nope"
	DatabaseNameTest                         = "../../DBTestAccount"
	ErrorOpenConnection                      = "Error on open connection"
	ErrorOpenConnectionNotImplemented        = "Error on open connection not implemented"
	ErrorOpenConnectionServiceNil            = "Error on open connection, service nil"
	ErrorOpenConnectionNotImplementedService = "Error on open connection not implemented, service is not nil"
)

func init() {
	os.Setenv(db.EnvDatabaseName, DatabaseNameTest)
}

func TestDbOpenSuccess(t *testing.T) {
	t.Run("Teste Abrir conexão com o banco de dados", func(t *testing.T) {
		srv, err := db.Open(db.Scribble)
		if err != nil {
			t.Error(ErrorOpenConnection, err)
		}

		if srv == nil {
			t.Error(ErrorOpenConnectionServiceNil)
		}
	})
}

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
