package model

import "time"

// Constantes para trabalhar com o model accounts.
const (
	AccountsTablename = "accounts"
)

// Account - Estrutura da entidade `account`
type Account struct {
	ID        string     `json:"id"`
	Cpf       string     `json:"cpf"`
	Name      string     `json:"name"`
	Secret    string     `json:"secret"`
	Balance   float64    `json:"balance"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// AccountBalance - Estrutura utilizada para serializar o balance da entidade `account`
type AccountBalance struct {
	Balance float64 `json:"balance"`
}

// AccountStorage - Interface que define as assinaturas para o storage de accounts.
type AccountStorage interface {
	GetAllAccounts() ([]Account, error)
	GetAccount(id string) (*Account, error)
	Insert(account Account) (*Account, error)
	Update(account Account) (*Account, error)
	GetAccountByCPF(cpf string) (*Account, error)
}
