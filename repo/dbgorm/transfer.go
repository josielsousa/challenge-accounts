package dbgorm

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/josielsousa/challenge-accounts/repo/model"
)

// TransferStorage - Assinatura para o storage de transfer.
type TransferStorage struct {
	db *gorm.DB
}

func (transferGorm) TableName() string {
	return model.TransferTablename
}

// Estrutura necessária para que o gorm realize o mapeamento das colunas com os tipos de dados.
type transferGorm struct {
	ID                   string  `gorm:"type:uuid;primary_key;"`
	AccountOriginID      string  `gorm:"type:varchar(255)"`
	AccountDestinationID string  `gorm:"type:varchar(255)"`
	Amount               float64 `gorm:"type:numeric(18,2)"`
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	DeletedAt            *time.Time `sql:"index"`
}

// NewTransferStorage - Inicializa o storage para transfers no banco de dados com Gorm.
func NewTransferStorage(db *gorm.DB) *TransferStorage {
	db.AutoMigrate(&model.Transfer{})

	return &TransferStorage{db: db}
}

// GetAllTransfers - Recupera todas as transfers da `account` informada.
func (s *TransferStorage) GetAllTransfers(accountID string) ([]model.Transfer, error) {
	transfers := make([]model.Transfer, 0)
	filter := &transferGorm{AccountOriginID: accountID}

	err := s.db.Table(model.TransferTablename).Where(filter).Find(&transfers).Error
	if err != nil {
		return nil, err
	}

	return transfers, err
}

// Insert - Insere uma nova transfer.
func (s *TransferStorage) Insert(transfer model.Transfer) (*model.Transfer, error) {
	if len(transfer.ID) <= 0 {
		transfer.ID = uuid.New().String()
	}

	err := s.db.Table(model.TransferTablename).Create(&transfer).Error
	if err != nil {
		return nil, err
	}

	return &transfer, err
}
