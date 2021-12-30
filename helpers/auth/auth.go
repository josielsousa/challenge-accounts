package auth

import "golang.org/x/crypto/bcrypt"

// Helper - Define a struct para o helper Auth.
type Helper struct{}

// NewHelper - Instância helper Auth.
func NewHelper() *Helper {
	return &Helper{}
}

// Hash - gera um novo hash para o secret da account.
func (h *Helper) Hash(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

// VerifySecret - verifica se a senha informada é a mesma salva na account.
func (h *Helper) VerifySecret(hashedSecret, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSecret), []byte(secret))
}
