package auth

import "golang.org/x/crypto/bcrypt"

//Helper - Define a struct para o helper HTTP.
type Helper struct {
}

//NewHelper - Instância helper HTTP.
func NewHelper() *Helper {
	return &Helper{}
}

//Hash - gera um novo hash para o secret da account.
func (s *Helper) Hash(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

//VerifySecret - verifica se a senha informada é a mesma salva na account.
func (s *Helper) VerifySecret(hashedSecret, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSecret), []byte(secret))
}
