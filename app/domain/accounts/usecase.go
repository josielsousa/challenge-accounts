package accounts

import (
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

type AccountInput struct {
	Name    string
	Balance int
	CPF     string
	Secret  string
}

type AccountOutput struct {
	ID        string
	Name      string
	Balance   int
	CPF       cpf.CPF
	CreatedAt time.Time
	UpdatedAt time.Time
}
