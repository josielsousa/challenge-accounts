package auth

import (
	"context"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/types"
)

type Usecase struct {
	R Repository
	H Hasher
	S Signer
}

//go:generate moq -rm -out usecase_mock.go . Signer Hasher Repository
type Signer interface {
	SignToken(id, username string) (types.Auth, error)
}

type Hasher interface {
	VerifySecret(hashedSecret, secret string) error
}

// Repository - Interface que define as assinaturas para o repository de accounts.
type Repository interface {
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
}

func NewUsecase(repo Repository, sig Signer, hasher Hasher) *Usecase {
	return &Usecase{
		R: repo,
		H: hasher,
		S: sig,
	}
}
