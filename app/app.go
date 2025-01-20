package app

import (
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/configuration"
	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
	pgacc "github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/accounts"
	pgtrf "github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/transfers"
)

type App struct {
	UA *accounts.Usecase
	UT *transfers.Usecase
}

func NewApp(cfg configuration.Config) (*App, error) {
	pgClient, err := postgres.NewClient(cfg.Postgres.URL())
	if err != nil {
		return nil, fmt.Errorf("on create postgres client: %w", err)
	}

	accRepo := pgacc.NewRepository(pgClient.Pool)
	accUC := accounts.NewUsecase(accRepo)
	trfUC := transfers.NewUsecase(pgtrf.NewRepository(pgClient.Pool), accRepo)

	return &App{
		UA: accUC,
		UT: trfUC,
	}, nil
}
