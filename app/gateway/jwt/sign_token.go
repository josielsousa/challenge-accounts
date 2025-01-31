package jwt

import (
	"fmt"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go/v4"

	"github.com/josielsousa/challenge-accounts/app/types"
)

const TTLToken = 5 * time.Minute

func (j Jwt) SignToken(id, username string) (string, error) {
	// JWT Claims - Payload contendo o CPF do usuário e a data de
	// expiração do token
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, &types.Claims{
		AccountID: id,
		Username:  username,
		ExpiresAt: j.C.Now().Add(TTLToken).Unix(),
	})

	tokenString, err := token.SignedString(j.appKey)
	if err != nil {
		return "", fmt.Errorf("on signed token: %w", err)
	}

	return tokenString, nil
}
