package db

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/josielsousa/challenge-accounts/repo/dbgorm"
)

// openScribble - Inicializa a conexão com o database scribble.
func openGorm() (service *Service, err error) {
	databaseName := os.Getenv(EnvDatabaseNameGorm)
	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		return nil, err
	}

	return getServicesGorm(db), nil
}

//getServicesGorm - Retorna as implementações de storage para o gorm.
func getServicesGorm(db *gorm.DB) *Service {
	return &Service{
		Account: dbgorm.NewAccountStorage(db),
	}
}
