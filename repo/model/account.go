package model

//Constantes para trabalhar com o model accounts.
const (
	Tablename = "accounts"
)

//Account - Estrutura da entidade `account`
type Account struct {
	ID        string      `json:"id, omitempty"`
	Cpf       string      `json:"cpf, omitempty"`
	Name      string      `json:"name, omitempty"`
	Secret    string      `json:"secret, omitempty"`
	Ballance  float64     `json:"ballance, omitempty"`
	CreatedAt interface{} `json:"created_at, omitempty"`
}

//AccountStorage - Interface que define as assinaturas para o storage de accounts.
type AccountStorage interface {
	GetAllAccounts() ([]string, error)
	GetAccount(id string) (*Account, error)
	Insert(account Account) (*Account, error)
	Update(account Account) (*Account, error)
}
