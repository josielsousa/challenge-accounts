package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/auth/helpers"
)

func TestAuthHelpers(t *testing.T) {
	t.Parallel()

	t.Run("gera e valida uma secret a partir de uma seed", func(t *testing.T) {
		t.Parallel()

		seed := "api-hash"
		helperAuth := helpers.NewHelper()

		secretHashed, err := helperAuth.Hash(seed)
		require.NoError(t, err)

		err = helperAuth.VerifySecret(string(secretHashed), seed)
		require.NoError(t, err)
	})
}
