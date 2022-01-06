package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidSecret = errors.New("invalid secret")
	ErrGenerateHash  = errors.New("unexpected error generating hash")
)

// GenHash - gera um novo hash para o secret da account.
func GenHash(secret string) (string, error) {
	hs, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrGenerateHash
	}

	return string(hs), nil
}

func CompareHashedAndSecret(hashedSecret, secret string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedSecret), []byte(secret))
	if err != nil {
		return ErrInvalidSecret
	}

	return nil
}
