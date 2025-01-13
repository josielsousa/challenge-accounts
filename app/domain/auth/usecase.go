package auth

import (
	"errors"

	"github.com/josielsousa/challenge-accounts/app/domain/auth/helpers"
	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

// Constantes utilizadas no serviço de autenticação.
//
//nolint:gosec
const (
	MaxTimeToExpiration = 5
	InfoTokenEmpty      = "Token vazio."
	InfoTokenExpired    = "Token expirado."
)

var (
	ErrTokenInvalid          = errors.New("token inválido")
	ErrTokenMalformed        = errors.New("token is malformed")
	ErrTokenSignatureInvalid = errors.New("token signature is invalid")
	ErrSignatureKeyInvalid   = errors.New("a chave de assinatura do token é inválida")
)

// JWT string chave utilizada para geração do token.
var jwtKey = []byte("api-challenge-accounts")

type Usecase struct {
	R accE.Repository
	H *helpers.Helper
}

func NewUsecase(repo accE.Repository) *Usecase {
	return &Usecase{
		R: repo,
		H: helpers.NewHelper(),
	}
}
