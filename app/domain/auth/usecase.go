package auth

import (
	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/types"
)

type Usecase struct {
	R accE.Repository
	H Hasher
	S Signer
}

//go:generate moq -rm -out usecase_mock.go . Signer Hasher
type Signer interface {
	SignToken(id, username string) (types.Auth, error)
}

type Hasher interface {
	VerifySecret(hashedSecret, secret string) error
}

func NewUsecase(repo accE.Repository, sig Signer, hasher Hasher) *Usecase {
	return &Usecase{
		R: repo,
		H: hasher,
		S: sig,
	}
}
