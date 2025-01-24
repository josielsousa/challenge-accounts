package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/auth"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/handler"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/middleware"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"

	// import swag docs
	_ "github.com/josielsousa/challenge-accounts/docs/swag"
)

type API struct {
	accUC  *accounts.Usecase
	trfUC  *transfers.Usecase
	authUC *auth.Usecase

	signer  *jwt.Jwt
	Handler http.Handler
}

//	@title			Challenge Accounts API
//	@version		1.0
//	@description	Implementação de API para o desafio de backend.
//	@description	A API é responsável por gerenciar contas e transferências
//	@description	entre contas.

//	@contact.name	Josiel Sousa
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api/v1/challenge-accounts

//	@securityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization

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
	)

	router.Route("/api/v1/challenge-accounts", func(baseRouter chi.Router) {
		baseRouter.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("/api/v1/challenge-accounts/swagger/doc.json"),
		))

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

		baseRouter.Route("/accounts", func(baseRouter chi.Router) {
			handler.RegisterAccountsHandlers(accUC, baseRouter)
			handler.RegisterTransfersHandlers(trfUC, signer, baseRouter)
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
