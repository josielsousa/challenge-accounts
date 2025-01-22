package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/auth"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/handler"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/middleware"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
)

type API struct {
	accUC  *accounts.Usecase
	trfUC  *transfers.Usecase
	authUC *auth.Usecase

	signer  *jwt.Jwt
	Handler http.Handler
}

func NewAPI(
	accUC *accounts.Usecase,
	trfUC *transfers.Usecase,
	authUC *auth.Usecase,
	signer *jwt.Jwt,
) *API {
	router := chi.NewRouter()

	router.Use(
		middleware.CORS(),
		middleware.CleanPath,
		middleware.StripSlashes,
		middleware.Recoverer,
		middleware.SetContentTypeJSON,
	)

	router.Route("/api/v1/challenge-accounts", func(baseRouter chi.Router) {
		baseRouter.Get("/healthcheck", func(rw http.ResponseWriter, _ *http.Request) {
			rw.WriteHeader(http.StatusOK)
		})

		baseRouter.Get("/", func(rw http.ResponseWriter, _ *http.Request) {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("<h1>Desafio técnico accounts.</h1>")) //nolint:errcheck
		})

		baseRouter.Route("/auth", func(authRouter chi.Router) {
			handler.RegisterAuthHandlers(authUC, authRouter)
		})

		baseRouter.Route("/accounts", func(accRouter chi.Router) {
			handler.RegisterAccountsHandlers(accUC, accRouter)
		})

		baseRouter.Route("/transfers", func(trfRouter chi.Router) {
			handler.RegisterTransfersHandlers(trfUC, signer, trfRouter)
		})
	})

	return &API{
		accUC:   accUC,
		trfUC:   trfUC,
		authUC:  authUC,
		signer:  signer,
		Handler: router,
	}
}
