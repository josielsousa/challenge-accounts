package handler

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
)

type (
	AccountBalanceResponse struct {
		// Balance - saldo atual da conta.
		Balance int `json:"balance"`
	} //	@name	GetAccountBalanceResponse
)

// GetAccountBalance godoc
//
//	@Summary		Retorna o saldo da conta.
//	@Description	Endpoint utilizado consultar o saldo da conta.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account_id	path		string						true	"Identificador da conta de origem"
//	@Success		200			{object}	GetAccountBalanceResponse	"Saldo da conta."
//	@Failure		400			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse	"Conta não encontrada."
//	@Failure		500			{object}	ErrorResponse
//	@Router			/api/v1/challenge-accounts/accounts/{account_id}/balance [get]
//
// TODO: add authorization here.
func (h Handler) GetAccountBalance(req *http.Request) *response.Response {
	id := chi.URLParam(req, "account_id")

	if len(strings.TrimSpace(id)) == 0 {
		return response.BadRequest(erring.ErrEmptyAccountID)
	}

	balance, err := h.accUC.GetAccountBalance(req.Context(), id)
	if err != nil {
		return response.AppError(err)
	}

	return response.Ok(AccountBalanceResponse{
		Balance: balance,
	})
}
