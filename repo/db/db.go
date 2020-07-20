package db

import (
	"errors"

	"github.com/josielsousa/challenge-accounts/repo/model"
)

// Constantes para o pacote repo.
const (
	Gorm                     = "gorm"
	ErrorDataBaseTypeInvalid = "Database escolhido não existe."
)

// Service - estrutura com todos os serviços disponíveis.
type Service struct {
	Account model.AccountStorage
}

//Open - Abre a conexão com o banco de dados.
func Open(dataBase string) (*Service, error) {
	switch dataBase {
	case Gorm:
		return openGorm()
	default:
		return nil, errors.New(ErrorDataBaseTypeInvalid)
	}
}
