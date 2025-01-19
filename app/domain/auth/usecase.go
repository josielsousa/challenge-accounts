package auth

import (
	"github.com/josielsousa/challenge-accounts/app/domain/auth/helpers"
	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/types"
)

type Usecase struct {
	R      accE.Repository
	H      *helpers.Helper
	S      Signer
	appKey []byte
}

//go:generate moq -rm -out usecase_mock.go . Signer
type Signer interface {
	SignToken(id string, username string) (types.Auth, error)
}

func NewUsecase(repo accE.Repository, sig Signer) *Usecase {
	return &Usecase{
		R: repo,
		H: helpers.NewHelper(),
		S: sig,

		// JWT string chave utilizada para geração do token.
		appKey: []byte("api-challenge-accounts"),
	}
}
