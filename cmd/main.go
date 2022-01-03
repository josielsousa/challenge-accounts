package main

import (
	"github.com/sirupsen/logrus"

	"github.com/josielsousa/challenge-accounts/app/configuration"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
)

//import (
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
	logger := logrus.New()
	logger.Info("api inicializando...")

	cfg, err := configuration.LoadConfig()
	if err != nil {
		logger.WithError(err).Error("on loading configuration")
	}

	log := logger.WithField("app", cfg.API.AppName)

	dbPool, err := postgres.ConnectPoolWithMigrations(cfg.Postgres.URL(), log, postgres.LogLevelError)
	if err != nil {
		logger.WithError(err).Error("on connect with database")
	}

	defer dbPool.Close()
	logger.Info("api avaiable...")
}
