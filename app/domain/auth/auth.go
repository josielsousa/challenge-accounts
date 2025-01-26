package auth

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

type (
	// SiginInput - estrutura utilizada para realizar a autenticação.
	SiginInput struct {
		Cpf    string
		Secret string
	}

	// SiginOutput - estrutura utilizada retornar o token da autenticação.
	SiginOutput struct {
		Token string
	}
)

func (u Usecase) Signin(ctx context.Context, input SiginInput) (SiginOutput, error) {
	const op = `auth.Signin`

	acc, err := u.R.GetByCPF(ctx, input.Cpf)
	if err != nil {
		return SiginOutput{}, fmt.Errorf(
			"%s-> %s: %w",
			op,
			"on get account by cpf",
			err,
		)
	}

	if len(acc.ID) == 0 {
		return SiginOutput{}, erring.ErrAccountNotFound
	}

	// Verifica se o secret informado na autenticação, é o mesmo armazenado na `account`.
	err = u.H.VerifySecret(acc.Secret.Value(), input.Secret)
	if err != nil {
		return SiginOutput{}, erring.ErrUnauthorized
	}

	// Tempo de expiração do token
	authToken, err := u.S.SignToken(acc.ID, acc.CPF.Value())
	if err != nil {
		return SiginOutput{}, erring.ErrUnexpected
	}

	return SiginOutput{Token: authToken}, nil
}
