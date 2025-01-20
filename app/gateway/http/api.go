package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/http/middleware"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
	"github.com/josielsousa/challenge-accounts/types"
)

//go:generate moq -rm -out api_mock.go . accUsecase authUsecase trfUsecase
type accUsecase interface {
	Create(ctx context.Context, input accounts.AccountInput) error
	GetAccountBalance(ctx context.Context, accountID string) (int, error)
	GetAllAccounts(ctx context.Context) ([]accounts.AccountOutput, error)
}

type trfUsecase interface {
	DoTransfer(ctx context.Context, input transfers.TransferInput) error
	ListTransfersAccount(ctx context.Context, accOriginID string) ([]transfers.TransferOutput, error)
}

type authUsecase interface {
	Signin(ctx context.Context, credential types.Credentials) (types.Auth, error)
}

type API struct {
	accUC   *accUsecase
	authUC  *authUsecase
	handler http.Handler
	signer  *jwt.Jwt
	trfUC   *trfUsecase
}

func NewAPI(
	accUC *accUsecase,
	trfUC *trfUsecase,
	authUC *authUsecase,
	signer *jwt.Jwt,
) *API {
	router := chi.NewRouter()

	router.Use(
		middleware.CORS(),
		middleware.CleanPath,
		middleware.StripSlashes,
		middleware.Recoverer,
	)
	router.Get("/healthcheck", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	// TODO: add all routers

	return &API{
		accUC:   accUC,
		trfUC:   trfUC,
		authUC:  authUC,
		signer:  signer,
		handler: router,
	}
}
