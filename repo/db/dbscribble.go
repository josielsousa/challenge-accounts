package db

import (
	"os"

	"github.com/josielsousa/challenge-accounts/repo/dbscribble"
	scribble "github.com/nanobox-io/golang-scribble"
)

// openScribble - Inicializa a conexão com o database scribble.
func openScribble() (service *Service, err error) {
	databaseName := os.Getenv(EnvDatabaseName)
	db, err := scribble.New(databaseName, nil)
	if err != nil {
		return nil, err
	}

	return getServicesScribble(db), nil
}

//getServicesScribble - Retorna as implementações de storage para o scribble.
func getServicesScribble(db *scribble.Driver) *Service {
	return &Service{
		Account: dbscribble.NewAccountStorage(db),
	}
}
