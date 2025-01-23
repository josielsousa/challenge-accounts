package accounts

import (
	"fmt"
	"os"
	"testing"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestMain(m *testing.M) {
	teardown, err := pgtest.StartupNewPool()
	if err != nil {
		panic(fmt.Errorf("on startup pgtest container: %w", err))
	}

	code := m.Run()

	teardown()

	os.Exit(code)
}
