package handler

import (
	"net/http"
	"time"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
)

type (
	ListAccountResponse struct {
		// Data - lista de contas.
		Data []AccountResponse `json:"data"`
		// Succes - indica se a operação foi bem sucedida.
		Succes bool `json:"success"`
	} //	@name	ListAccountResponse
)

// ListAccounts godoc
//
//	@Summary		Lista as contas.
//	@Description	Endpoint utilizado listar todas as contas.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ListAccountResponse	"Lista todas as contas"
//	@Failure		400	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/v1/challenge-accounts/accounts [get]
func (h Handler) ListAccounts(req *http.Request) *response.Response {
	out, err := h.accUC.GetAllAccounts(req.Context())
	if err != nil {
		return response.AppError(err)
	}

	accs := make([]AccountResponse, 0, len(out))

	for _, acc := range out {
		accs = append(accs, AccountResponse{
			ID:        acc.ID,
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
