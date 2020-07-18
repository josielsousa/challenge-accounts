package main

import (
	"github.com/josielsousa/challenge-accounts/providers/http"
	"github.com/josielsousa/challenge-accounts/providers/log"
)

func main() {
	logger := log.NewLogger()
	logger.Info("API inicializando...")

	router := http.NewRouter(logger)
	router.ServeHTTP()
}
