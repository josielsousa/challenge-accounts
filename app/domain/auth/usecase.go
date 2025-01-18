package auth

import (
	"errors"
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

var (
	ErrEmptyToken           = errors.New("token is empty")
	ErrExpiredToken         = errors.New("o token está expirado")
	ErrInvalidToken         = errors.New("token inválido")
	ErrMalformedToken       = errors.New("token mal formado")
	ErrParseTokenWithClaims = errors.New("o token não pode ser analisado")
	ErrSignatureKeyInvalid  = errors.New("a chave de assinatura do token é inválida")
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
