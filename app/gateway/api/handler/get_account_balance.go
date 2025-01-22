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
		Balance int `json:"balance"`
	}
)

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
