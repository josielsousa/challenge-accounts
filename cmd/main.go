package main

import (
	"log/slog"
	"os"

	"github.com/josielsousa/challenge-accounts/app/configuration"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
)

// import (
// "github.com/josielsousa/challenge-accounts/providers/http"
// "github.com/josielsousa/challenge-accounts/providers/log"
// "github.com/josielsousa/challenge-accounts/repo/db"
// "github.com/josielsousa/challenge-accounts/service"
// "github.com/josielsousa/challenge-accounts/types"
// _ "github.com/mattn/go-sqlite3"
//)
// func main() {
// 	logger := log.NewLogger()
// 	logger.Info("API inicializando...")

// 	stg, err := db.Open(db.Gorm)
// 	if err != nil {
// 		logger.Error(types.ErrorOpenConnection, err)
// 		return
// 	}

// 	defer func() {
// 		stg.Close()
// 	}()

// 	srvAuth := service.NewAuthService(stg.Account, logger)
// 	srvTransfer := service.NewTransferService(stg, logger)
// 	srvAccount := service.NewAccountService(stg.Account, logger)

// 	router := http.NewRouter(srvAuth, srvAccount, srvTransfer, logger)
// 	router.ServeHTTP()
// }

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("api inicializando...")

	cfg, err := configuration.LoadConfig()
	if err != nil {
		logger.Error("on loading configuration", slog.Any("error", err))
	}

	// JWT string chave utilizada para geração do token.
	// jwt := jwt.New([]byte("api-challenge-accounts"))

	dbPool, err := postgres.ConnectPoolWithMigrations(cfg.Postgres.URL())
	if err != nil {
		logger.Error("on connect to pool", slog.Any("error", err))
	}

	defer dbPool.Close()
	logger.Info("api available...")
}
