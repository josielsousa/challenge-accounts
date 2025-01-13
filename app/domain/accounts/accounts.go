package accounts

import (
	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

type Account struct {
	accRepo accE.Repository
}

func NewUsecase(accRepo accE.Repository) *Account {
	return &Account{accRepo: accRepo}
}
