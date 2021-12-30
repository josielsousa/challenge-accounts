package db

import (
	"github.com/jinzhu/gorm"

	"github.com/josielsousa/challenge-accounts/repo/dbgorm"
	"github.com/josielsousa/challenge-accounts/types"
)

// openScribble - Inicializa a conexão com o database scribble.
func openGorm() (service *Service, err error) {
	db, err := gorm.Open("sqlite3", types.DatabaseName)
	if err != nil {
		return nil, err
	}

	return getServicesGorm(db), nil
}

// getServicesGorm - Inicializa uma transação no banco de dados.
func (s *Service) openGormTransaction() *Service {
	tx := s.conn.(*gorm.DB).Begin()
	return getServicesGorm(tx)
}

// getServicesGorm - Retorna as implementações de storage para o gorm.
func getServicesGorm(db *gorm.DB) *Service {
	return &Service{
		conn:     db,
		connType: Gorm,
		Account:  dbgorm.NewAccountStorage(db),
		Transfer: dbgorm.NewTransferStorage(db),
	}
}
