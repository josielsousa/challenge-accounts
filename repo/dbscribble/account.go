package dbscribble

import (
	"encoding/json"

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
func (s *AccountStorage) GetAllAccounts() ([]model.Account, error) {
	accountsDb, err := s.db.ReadAll(model.AccountsTablename)
	if err != nil {
		return nil, err
	}

	var accountsJSON []model.Account
	for _, account := range accountsDb {
		if len(account) > 0 {
			accountJSON := &model.Account{}
			err = json.Unmarshal([]byte(account), accountJSON)
			if err != nil {
				return nil, err
			}

			accountsJSON = append(accountsJSON, *accountJSON)
		}
	}

	return accountsJSON, err
}

//GetAccount - Recupera uma account conforme o `id` informado.
func (s *AccountStorage) GetAccount(id string) (*model.Account, error) {
	account := &model.Account{}
	err := s.db.Read(model.AccountsTablename, id, account)
	if err != nil {
		return nil, err
	}

	return account, err
}

//Insert - Insere uma nova account.
func (s *AccountStorage) Insert(account model.Account) (*model.Account, error) {
	err := s.db.Write(model.AccountsTablename, account.ID, account)
	if err != nil {
		return nil, err
	}

	return &account, err
}

//Update - Atualiza a account informada.
func (s *AccountStorage) Update(account model.Account) (*model.Account, error) {
	err := s.db.Write(model.AccountsTablename, account.ID, account)
	if err != nil {
		return nil, err
	}

	return &account, err
}
