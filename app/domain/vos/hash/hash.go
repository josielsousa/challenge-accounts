package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidSecret = errors.New("invalid secret")
	ErrGenerateHash  = errors.New("unexpected error generating hash")
)

type Hash struct {
	value string
}

func (h Hash) String() string {
	return h.value
}

func NewHash(secret string) (Hash, error) {
	h := Hash{}
	hs, err := h.hash(secret)
	if err != nil {
		return Hash{}, ErrGenerateHash
	}

	h.value = string(hs)
	return h, nil
}

// Hash - gera um novo hash para o secret da account.
func (h Hash) hash(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

// CompareHashAndSecret - verifica se a senha informada é a mesma salva na account.
func (h Hash) CompareHashAndSecret(secret string) error {
	err := bcrypt.CompareHashAndPassword([]byte(h.value), []byte(secret))
	if err != nil {
		return ErrInvalidSecret
	}

	return nil
}
