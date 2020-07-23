package db

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/josielsousa/challenge-accounts/repo/model"
)

// Constantes para o pacote repo.
const (
	Gorm                     = "gorm"
	ErrorDataBaseTypeInvalid = "Database escolhido não existe."
)

// Service - estrutura com todos os serviços disponíveis.
type Service struct {
	conn     interface{}
	connType string
	Account  model.AccountStorage
	Transfer model.TransferStorage
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

//Close - Fecha a conexão com o banco de dados.
func (s *Service) Close() {
	switch s.connType {
	case Gorm:
		if s.conn != nil {
			s.conn.(*gorm.DB).Close()
		}

		s.conn = nil
	}
}

//BeginTransaction - Inicia a transação no banco de dados.
func (s *Service) BeginTransaction() *Service {
	switch s.connType {
	case Gorm:
		return s.openGormTransaction()
	}

	return nil
}

//Rollback - Realiza o rollback da transação no banco de dados.
func (s *Service) Rollback() {
	switch s.connType {
	case Gorm:
		s.conn.(*gorm.DB).Rollback()
	}
}

//Commit - Realiza o commit da transação no banco de dados.
func (s *Service) Commit() {
	switch s.connType {
	case Gorm:
		s.conn.(*gorm.DB).Commit()
	}
}
