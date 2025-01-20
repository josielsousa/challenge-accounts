package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/middleware"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
	"github.com/josielsousa/challenge-accounts/types"
)

//go:generate moq -rm -out handler_mock.go . accUsecase authUsecase trfUsecase
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

type Handler struct {
	accUC  accUsecase
	authUC authUsecase
	trfUC  trfUsecase
}

func RegisterAuthHandlers(authUC authUsecase, router chi.Router) {
	handler := &Handler{authUC: authUC}

	router.Post("/login", rest.Handler(handler.Login))
}

func RegisterAccountsHandlers(accUC accUsecase, router chi.Router) {
	handler := &Handler{accUC: accUC}

	router.Post("/", rest.Handler(handler.NoContent))
}

func RegisterTransfersHandlers(trfUC trfUsecase, signer *jwt.Jwt, router chi.Router) {
	handler := &Handler{trfUC: trfUC}

	// List all transfers
	router.Get("/", middleware.Authorize(signer, rest.Handler(handler.NoContent)))

	// Create a transfer
	router.Post("/", middleware.Authorize(signer, rest.Handler(handler.NoContent)))
}

func (Handler) NoContent(_ *http.Request) *response.Response {
	return response.NoContent()
}
