package model

import "time"

//Constantes para trabalhar com o model accounts.
const (
	AccountsTablename = "accounts"
)

//Account - Estrutura da entidade `account`
type Account struct {
	ID        string     `json:"id,omitempty"`
	Cpf       string     `json:"cpf,omitempty"`
	Name      string     `json:"name,omitempty"`
	Secret    string     `json:"secret,omitempty"`
	Ballance  float64    `json:"ballance,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

//AccountStorage - Interface que define as assinaturas para o storage de accounts.
type AccountStorage interface {
	GetAllAccounts() ([]Account, error)
	GetAccount(id string) (*Account, error)
	Insert(account Account) (*Account, error)
	Update(account Account) (*Account, error)
}
