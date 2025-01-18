package auth

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	"github.com/josielsousa/challenge-accounts/types"
)

func (u Usecase) getAppKey(_ *jwt.Token) (interface{}, error) {
	return u.appKey, nil
}

func (u Usecase) Authorize(token string) (types.Claims, error) {
	claims := &types.Claims{}

	if len(strings.TrimSpace(token)) == 0 || token == "null" {
		return types.Claims{}, ErrEmptyToken
	}

	jwtToken, err := jwt.ParseWithClaims(token, claims, u.getAppKey)
	if err != nil {
		return types.Claims{}, ErrParseTokenWithClaims
	}

	if !jwtToken.Valid {
		return types.Claims{}, ErrInvalidToken
	}

	minutesElapsedLastAuth := time.Until(time.Unix(claims.ExpiresAt, 0))

	if minutesElapsedLastAuth <= 0 {
		return types.Claims{}, ErrExpiredToken
	}

	return types.Claims{}, nil
}
