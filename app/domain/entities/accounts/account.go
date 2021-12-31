package accounts

import (
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

// Account - Estrutura da entidade `account`
type Account struct {
	ID        string
	Name      string
	Secret    string
	Balance   int
	CPF       cpf.CPF
	CreatedAt time.Time
	UpdatedAt time.Time
}
