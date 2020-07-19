package dbgorm

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/josielsousa/challenge-accounts/repo/model"
)

//AccountStorage - Assinatura para o storage de account.
type AccountStorage struct {
	db *gorm.DB
}

func (accountGorm) TableName() string {
	return model.AccountsTablename
}

// Estrutura necessária para que o gorm realize o mapeamento das colunas com os tipos de dados.
type accountGorm struct {
	ID        string  `gorm:"type:uuid;primary_key;"`
	Cpf       string  `gorm:"type:varchar(255)"`
	Name      string  `gorm:"type:varchar(255)"`
	Secret    string  `gorm:"type:varchar(255)"`
	Ballance  float64 `gorm:"type:numeric(18,2)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//NewAccountStorage - Inicializa o storage para accounts no banco de dados com Gorm.
func NewAccountStorage(db *gorm.DB) *AccountStorage {
	db.AutoMigrate(&accountGorm{})

	return &AccountStorage{db: db}
}

//GetAllAccounts - Recupera todas as accounts.
func (s *AccountStorage) GetAllAccounts() ([]model.Account, error) {
	accounts := make([]model.Account, 0)
	err := s.db.Table(model.AccountsTablename).Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, err
}

//GetAccount - Recupera uma account conforme o `id` informado.
func (s *AccountStorage) GetAccount(id string) (*model.Account, error) {
	account := &model.Account{}
	err := s.db.Table(model.AccountsTablename).Where(map[string]interface{}{"id": id}).First(&account).Error
	if err != nil {
		return nil, err
	}

	return account, err
}

//Insert - Insere uma nova account.
func (s *AccountStorage) Insert(account model.Account) (*model.Account, error) {
	if len(account.ID) <= 0 {
		account.ID = uuid.New().String()
	}

	err := s.db.Table(model.AccountsTablename).Create(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, err
}

//Update - Atualiza a account informada.
func (s *AccountStorage) Update(account model.Account) (*model.Account, error) {
	err := s.db.Table(model.AccountsTablename).Model(&accountGorm{}).Updates(account).Error
	if err != nil {
		return nil, err
	}

	return &account, err
}
