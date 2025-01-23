package jwt

import (
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go/v4"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/types"
)

func (j Jwt) getAppKey(_ *jwtgo.Token) (interface{}, error) {
	return j.appKey, nil
}

func (j Jwt) Authorize(token string) (types.Claims, error) {
	claims := &types.Claims{}

	if len(strings.TrimSpace(token)) == 0 || token == "null" {
		return types.Claims{}, erring.ErrEmptyToken
	}

	jwtToken, err := jwtgo.ParseWithClaims(token, claims, j.getAppKey)
	if err != nil {
		return types.Claims{}, erring.ErrParseTokenWithClaims
	}

	if !jwtToken.Valid {
		return types.Claims{}, erring.ErrInvalidToken
	}

	minutesElapsedLastAuth := j.C.Until(time.Unix(claims.ExpiresAt, 0))

	if minutesElapsedLastAuth <= 0 {
		return types.Claims{}, erring.ErrExpiredToken
	}

	return *claims, nil
}
