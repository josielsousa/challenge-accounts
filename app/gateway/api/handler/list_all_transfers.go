package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/types"
)

type (
	ListTransfersResponse struct {
		Data   []TransferResponse `json:"data"`
		Succes bool               `json:"success"`
	}
)

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
