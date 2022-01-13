package usecase

import (
	accUC "github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

var _ accUC.Usecase = Account{}

type Account struct {
	accRepo accounts.Repository
}

func NewUsecase(accRepo accounts.Repository) *Account {
	return &Account{accRepo: accRepo}
}
