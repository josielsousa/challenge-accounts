package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/validator"
	"github.com/josielsousa/challenge-accounts/types"
)

type (
	DoTransferRequest struct {
		AccountDestinationID string `json:"account_destination_id" validate:"required"`
		Amount               int    `json:"amount"                 validate:"required"`
	}

	DoTransferResponse struct {
		ID                   string `json:"id"`
		AccountOriginID      string `json:"account_origin_id"`
		AccountDestinationID string `json:"account_destination_id"`
		Amount               int    `json:"amount"`
		CreatedAt            string `json:"created_at"`
	}
)

func (r DoTransferRequest) Validate() error {
	globalValidator := validator.GetGlobalValidator()

	if err := globalValidator.ValidateStructModel(r); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r DoTransferRequest) ToTransferInput(accountOriginID string) transfers.TransferInput {
	return transfers.TransferInput{
		AccountOriginID:      accountOriginID,
		AccountDestinationID: r.AccountDestinationID,
		Amount:               r.Amount,
	}
}

func (h Handler) DoTransfer(req *http.Request) *response.Response {
	id := chi.URLParam(req, "account_id")

	logger := slog.Default()

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

	var request DoTransferRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return response.BadRequest(err)
	}

	input := request.ToTransferInput(id)

	logger.Info("on do transfer", slog.Any("input", input))

	out, err := h.trfUC.DoTransfer(ctx, input)
	if err != nil {
		return response.AppError(err)
	}

	return response.Ok(DoTransferResponse{
		ID:                   out.ID,
		AccountOriginID:      out.AccountOriginID,
		AccountDestinationID: out.AccountDestinationID,
		Amount:               out.Amount,
		CreatedAt:            out.CreatedAt.Format(time.RFC3339),
	})
}
