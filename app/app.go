package app

import (
	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/auth"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
	pgacc "github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/accounts"
	pgtrf "github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/hasher"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
)

type App struct {
	UA   *accounts.Usecase
	UT   *transfers.Usecase
	UAT  *auth.Usecase
	Name string
}

func NewApp(
	pgClient *postgres.Client,
	signer *jwt.Jwt,
	hasher *hasher.Helper,
) *App {
	accRepo := pgacc.NewRepository(pgClient.Pool)

	accUC := accounts.NewUsecase(accRepo)

	trfUC := transfers.NewUsecase(pgtrf.NewRepository(pgClient.Pool), accRepo)

	authUC := auth.NewUsecase(accRepo, signer, hasher)

	return &App{
		UA:   accUC,
		UT:   trfUC,
		UAT:  authUC,
		Name: "Challenge Accounts",
	}
}
