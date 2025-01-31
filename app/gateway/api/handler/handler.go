package handler

import (
	"context"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/auth"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/middleware"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
)

//go:generate moq -rm -out handler_mock.go . accUsecase authUsecase trfUsecase
type accUsecase interface {
	Create(
		ctx context.Context, input accounts.AccountInput,
	) (accounts.AccountOutput, error)
	GetAccountBalance(ctx context.Context, accountID string) (int, error)
	GetAllAccounts(ctx context.Context) ([]accounts.AccountOutput, error)
}

type trfUsecase interface {
	DoTransfer(
		ctx context.Context, input transfers.TransferInput,
	) (transfers.TransferOutput, error)
	ListTransfersAccount(
		ctx context.Context, accOriginID string,
	) ([]transfers.TransferOutput, error)
}

type authUsecase interface {
	Signin(
		ctx context.Context, credential auth.SiginInput,
	) (auth.SiginOutput, error)
}

type Handler struct {
	accUC  accUsecase
	authUC authUsecase
	trfUC  trfUsecase
}

func RegisterAuthHandlers(authUC authUsecase, router chi.Router) {
	handler := &Handler{authUC: authUC}

	router.Post("/signin", rest.Handler(handler.Signin))
}

func RegisterAccountsHandlers(accUC accUsecase, router chi.Router) {
	handler := &Handler{accUC: accUC}

	router.Post("/", rest.Handler(handler.CreateAccount))
	router.Get("/", rest.Handler(handler.ListAccounts))

	router.Get("/{account_id}/balance",
		rest.Handler(handler.GetAccountBalance),
	)
}

func RegisterTransfersHandlers(
	trfUC trfUsecase,
	signer *jwt.Jwt,
	router chi.Router,
) {
	handler := &Handler{trfUC: trfUC}

	// List all transfers
	router.Get("/{account_id}/transfers",
		middleware.Authorize(
			signer,
			rest.Handler(handler.ListTransfers)),
	)

	// Create a transfer
	router.Post("/{account_id}/transfers",
		middleware.Authorize(
			signer, rest.Handler(handler.DoTransfer)),
	)
}
