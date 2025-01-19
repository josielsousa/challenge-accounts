package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
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
	authToken, err := u.SignToken(acc)
	if err != nil {
		return types.Auth{}, erring.ErrUnexpected
	}

	return authToken, nil
}

func (u Usecase) SignToken(acc accE.Account) (types.Auth, error) {
	// JWT Claims - Payload contendo o CPF do usuário e a data de expiração do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.Claims{
		AccountID: acc.ID,
		Username:  acc.CPF.Value(),
		ExpiresAt: time.Now().Add(TTLToken).Unix(),
	})

	tokenString, err := token.SignedString(u.appKey)
	if err != nil {
		return types.Auth{}, fmt.Errorf("on signed token: %w", err)
	}

	return types.Auth{Token: tokenString}, nil
}
