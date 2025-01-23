package auth

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/types"
)

func (u Usecase) Signin(ctx context.Context, credential types.Credentials) (types.Auth, error) {
	const op = `auth.Signin`

	acc, err := u.R.GetByCPF(ctx, credential.Cpf)
	if err != nil {
		return types.Auth{}, fmt.Errorf(
			"%s-> %s: %w",
			op,
			"on get account by cpf",
			err,
		)
	}

	if len(acc.ID) == 0 {
		return types.Auth{}, erring.ErrAccountNotFound
	}

	// Verifica se o secret informado na autenticação, é o mesmo armazenado na `account`.
	err = u.H.VerifySecret(acc.Secret.Value(), credential.Secret)
	if err != nil {
		return types.Auth{}, erring.ErrUnauthorized
	}

	// Tempo de expiração do token
	authToken, err := u.S.SignToken(acc.ID, acc.CPF.Value())
	if err != nil {
		return types.Auth{}, erring.ErrUnexpected
	}

	return authToken, nil
}
