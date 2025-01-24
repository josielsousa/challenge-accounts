package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/types"
)

type (
	ListTransfersResponse struct {
		// Data - lista de transferências.
		Data []TransferResponse `json:"data"`
		// Succes - indica se a operação foi bem sucedida.
		Succes bool `json:"success"`
	} //	@name	ListTransfersResponse
)

// ListTransfers godoc
//
//	@Summary		Lista as transferência da conta.
//	@Description	Endpoint utilizado para listar as transferências entre contas.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			account_id	path		string					true	"Identificador da conta de origem"
//	@Success		200			{object}	ListTransfersResponse	"Transferência realizada com sucesso."
//	@Failure		400			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse	"Conta não encontrada."
//	@Failure		422			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/api/v1/challenge-accounts/accounts/{account_id}/transfers [get]
func (h Handler) ListTransfers(req *http.Request) *response.Response {
	id := chi.URLParam(req, "account_id")

	if len(strings.TrimSpace(id)) == 0 {
		return response.BadRequest(erring.ErrEmptyAccountID)
	}

	ctx := req.Context()

	claims, ok := types.ClaimsFromContext(ctx)
	if !ok {
		return response.Unauthenticated()
	}

	if id != claims.AccountID {
		return response.Forbidden()
	}

	transfers, err := h.trfUC.ListTransfersAccount(ctx, id)
	if err != nil {
		return response.AppError(err)
	}

	out := make([]TransferResponse, 0, len(transfers))

	for _, trf := range transfers {
		out = append(out, TransferResponse{
			ID:                   trf.ID,
			AccountOriginID:      trf.AccountOriginID,
			AccountDestinationID: trf.AccountDestinationID,
			Amount:               trf.Amount,
			CreatedAt:            trf.CreatedAt.Format(time.RFC3339),
		})
	}

	return response.Ok(ListTransfersResponse{
		Data:   out,
		Succes: true,
	})
}
