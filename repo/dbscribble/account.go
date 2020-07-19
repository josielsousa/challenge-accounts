package dbscribble

import (
	"github.com/josielsousa/challenge-accounts/repo/model"
	scribble "github.com/nanobox-io/golang-scribble"
)

//AccountStorage - Assinatura para o storage de account.
type AccountStorage struct {
	db *scribble.Driver
}

//NewAccountStorage - Inicializa o storage para accounts no banco de dados Scribble.
func NewAccountStorage(db *scribble.Driver) *AccountStorage {
	return &AccountStorage{db: db}
}

//GetAllAccounts - Recupera todas as accounts.
func (s *AccountStorage) GetAllAccounts() ([]string, error) {
	accounts, err := s.db.ReadAll(model.Tablename)
	return accounts, err
}

//GetAccount - Recupera uma account conforme o `id` informado.
func (s *AccountStorage) GetAccount(id string) (*model.Account, error) {
	account := &model.Account{}
	err := s.db.Read(model.Tablename, id, account)
	return account, err
}

//Insert - Insere uma nova account.
func (s *AccountStorage) Insert(account model.Account) (*model.Account, error) {
	err := s.db.Write(model.Tablename, account.ID, account)
	return &account, err
}

//Update - Atualiza a account informada.
func (s *AccountStorage) Update(account model.Account) (*model.Account, error) {
	err := s.db.Write(model.Tablename, account.ID, account)
	return &account, err
}
