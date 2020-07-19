package db

import (
	"errors"

	"github.com/josielsousa/challenge-accounts/repo/model"
)

// Constantes para o pacote repo.
const (
	Scribble                 = "scribble"
	EnvDatabaseName          = "DATABASE_NAME"
	ErrorDataBaseTypeInvalid = "Database escolhido não existe."
)

// Service - estrutura com todos os serviços disponíveis.
type Service struct {
	Account model.AccountStorage
}

//Open - Abre a conexão com o banco de dados.
func Open(dataBase string) (*Service, error) {
	switch dataBase {
	case Scribble:
		return openScribble()
	default:
		return nil, errors.New(ErrorDataBaseTypeInvalid)
	}
}
