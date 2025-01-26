package types

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

type ClaimsKey struct{}

func ContextWithClaims(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, ClaimsKey{}, claims)
}

func ClaimsFromContext(ctx context.Context) (Claims, bool) {
	claims, ok := ctx.Value(ClaimsKey{}).(Claims)
	if !ok {
		return Claims{}, false
	}

	return claims, true
}

// Claims - Estrutura utilizada para trafegar as informações do token.
type Claims struct {
	Username  string `json:"username"`
	AccountID string `json:"account_id"`
	ExpiresAt int64  `json:"expires_at"`
	jwt.StandardClaims
}
