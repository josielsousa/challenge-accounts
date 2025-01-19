package auth

import (
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/auth/helpers"
	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

// Constantes utilizadas no serviço de autenticação.
//
//nolint:gosec
const (
	TTLToken         = 5 * time.Minute
	InfoTokenEmpty   = "Token vazio."
	InfoTokenExpired = "Token expirado."
)

type Usecase struct {
	R      accE.Repository
	H      *helpers.Helper
	appKey []byte
}

func NewUsecase(repo accE.Repository) *Usecase {
	return &Usecase{
		R: repo,
		H: helpers.NewHelper(),

		// JWT string chave utilizada para geração do token.
		appKey: []byte("api-challenge-accounts"),
	}
}
