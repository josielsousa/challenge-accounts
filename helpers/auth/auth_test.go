package auth_test

import (
	"testing"
	auth "github.com/josielsousa/challenge-accounts/helpers/auth"
)

func TestAuthHelpers(t *testing.T) {
	t.Run("Teste hash e validação do hash gerado", func(t *testing.T) {
		secret := "api-hash"
		helperAuth := auth.NewHelper()

		secretHashed, err := helperAuth.Hash(secret)
		if err != nil {
			t.Error("Error on hash secret: ", err)
			return
		}

		err = helperAuth.VerifySecret(string(secretHashed), secret)
		if err != nil {
			t.Error("Error on verify hashed secret: ", err)
			return
		}
	})
}
