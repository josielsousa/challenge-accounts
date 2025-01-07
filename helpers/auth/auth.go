package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Helper - Define a struct para o helper Auth.
type Helper struct{}

// NewHelper - Instância helper Auth.
func NewHelper() *Helper {
	return &Helper{}
}

// Hash - gera um novo hash para o secret da account.
func (h *Helper) Hash(secret string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generating hash: %w", err)
	}

	return hash, nil
}

// VerifySecret - verifica se a senha informada é a mesma salva na account.
func (h *Helper) VerifySecret(hashedSecret, secret string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedSecret), []byte(secret)); err != nil {
		return fmt.Errorf("comparing hash and secret: %w", err)
	}

	return nil
}
