package handler

import (
	"net/http"
	"time"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
)

type (
	ListAccountResponse struct {
		Data   []AccountResponse `json:"data"`
		Succes bool              `json:"success"`
	}
)

func (h Handler) ListAccounts(req *http.Request) *response.Response {
	out, err := h.accUC.GetAllAccounts(req.Context())
	if err != nil {
		return response.AppError(err)
	}

	accs := make([]AccountResponse, 0, len(out))

	for _, acc := range out {
		accs = append(accs, AccountResponse{
			Name:      acc.Name,
			Balance:   acc.Balance,
			CPF:       acc.CPF,
			CreatedAt: acc.CreatedAt.Format(time.RFC3339),
		})
	}

	return response.Ok(ListAccountResponse{
		Data:   accs,
		Succes: true,
	})
}
