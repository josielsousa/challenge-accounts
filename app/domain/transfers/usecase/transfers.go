package usecase

import (
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
	trfUC "github.com/josielsousa/challenge-accounts/app/domain/transfers"
)

var _ trfUC.Usecase = Transfer{}

type Transfer struct {
	repo    transfers.Repository
	accRepo accounts.Repository
}

func NewUsecase(trfRepo transfers.Repository, accRepo accounts.Repository) *Transfer {
	return &Transfer{
		repo:    trfRepo,
		accRepo: accRepo,
	}
}
