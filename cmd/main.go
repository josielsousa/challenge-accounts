package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/josielsousa/challenge-accounts/app"
	"github.com/josielsousa/challenge-accounts/app/configuration"
	"github.com/josielsousa/challenge-accounts/app/gateway/api"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
	"github.com/josielsousa/challenge-accounts/app/gateway/hasher"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
)

const (
	GracefulTimeout   = 5 * time.Second
	APIGeneralTimeout = 15 * time.Second
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("api inicializando...")

	ctx := context.Background()

	cfg, err := configuration.LoadConfig()
	if err != nil {
		logger.Error("on loading configuration", slog.Any("error", err))
	}

	pgClient, err := postgres.NewClient(cfg.Postgres.URL())
	if err != nil {
		logger.Error("on create configuration", slog.Any("error", err))

		panic(err)
	}

	// JWT string chave utilizada para geração do token.
	signer := jwt.New([]byte("api-challenge-accounts"))

	// Hasher é um helper usado para gerar e validar a secret.
	hasher := hasher.NewHelper()

	application := app.NewApp(pgClient, signer, hasher)

	challengeAPI := api.NewAPI(
		application.UA,
		application.UT,
		application.UAT,
		signer,
	)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      challengeAPI.Handler,
		WriteTimeout: APIGeneralTimeout,
		ReadTimeout:  APIGeneralTimeout,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
	}

	// Graceful Shutdown
	gracefulCtx, cancelGraceful := signal.NotifyContext(
		ctx,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	gracefulGroup, gracefulGroupCtx := errgroup.WithContext(gracefulCtx)

	gracefulGroup.Go(func() error {
		log.Printf("api available on server port: %s", server.Addr)

		return server.ListenAndServe()
	})

	//nolint:contextcheck
	gracefulGroup.Go(func() error {
		<-gracefulGroupCtx.Done()

		log.Printf("exit signal received terminating...")

		timeoutCtx, cancel := context.WithTimeout(context.Background(), GracefulTimeout)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			return fmt.Errorf("on stop server: %w", err)
		}

		pgClient.Pool.Close()

		return nil
	})

	cancelGraceful()
}
