package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/types"
)

// TODO: resolve no linters annotations.
//
//nolint:stylecheck
func (u Usecase) Auth(ctx context.Context, credential types.Credentials) (types.Auth, error) {
	const op = `auth.Auth`

	acc, err := u.R.GetByCPF(ctx, credential.Cpf)
	if err != nil {
		return types.Auth{}, fmt.Errorf("%s-> %s: %w", op, "on get account by cpf", err)
	}

	if len(acc.ID) == 0 {
		return types.Auth{}, errors.New(types.ErroAccountNotFound)
	}

	// Verifica se o secret informado na autenticação, é o mesmo armazenado na `account`.
	err = u.H.VerifySecret(acc.Secret.Value(), credential.Secret)
	if err != nil {
		return types.Auth{}, errors.New(types.ErrorUnauthorized)
	}

	// Tempo de expiração do token
	expirationTime := time.Now().Add(MaxTimeToExpiration * time.Minute)

	authToken, err := u.GetToken(acc, jwtKey, expirationTime)
	if err != nil {
		return types.Auth{}, errors.New(types.ErrorUnexpected)
	}

	return authToken, nil
}

func (u Usecase) GetToken(acc accE.Account, jwtKey []byte, expirationTime time.Time) (types.Auth, error) {
	// JWT Claims - Payload contendo o CPF do usuário e a data de expiração do token
	claims := &types.Claims{
		AccountID: acc.ID,
		Username:  acc.CPF.Value(),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return types.Auth{}, fmt.Errorf("on signed token: %w", err)
	}

	authToken := types.Auth{Token: tokenString}

	return authToken, nil
}
