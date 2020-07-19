package main

import (
	"github.com/josielsousa/challenge-accounts/providers/http"
	"github.com/josielsousa/challenge-accounts/providers/log"
	"github.com/josielsousa/challenge-accounts/repo/db"
	"github.com/josielsousa/challenge-accounts/service"
	"github.com/josielsousa/challenge-accounts/types"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger := log.NewLogger()
	logger.Info("API inicializando...")

	stg, err := db.Open(db.Gorm)
	if err != nil {
		logger.Error(types.ErrorOpenConnection, err)
	}

	srvAccount := service.NewAccountService(stg.Account, logger)

	router := http.NewRouter(srvAccount, logger)
	router.ServeHTTP()
}
