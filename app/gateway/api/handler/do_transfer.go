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
	"github.com/josielsousa/challenge-accounts/app/types"
)

type (
	DoTransferRequest struct {
		// AccountDestinationID - identificador da conta de destino.
		AccountDestinationID string `json:"account_destination_id" validate:"required"`
		// Amount - valor a ser transferido.
		Amount int `json:"amount" validate:"required"`
	} //	@name	DoTransferRequest

	TransferResponse struct {
		// ID - identificador único da transferência.
		ID string `json:"id"`
		// AccountOriginID - identificador da conta de origem.
		AccountOriginID string `json:"account_origin_id"`
		// AccountDestinationID - identificador da conta de destino.
		AccountDestinationID string `json:"account_destination_id"`
		// Amount - valor transferido.
		Amount int `json:"amount"`
		// CreatedAt - data de criação da transferência.
		CreatedAt string `json:"created_at"`
	} //	@name	TransferResponse
)

func (r DoTransferRequest) Validate() error {
	globalValidator := validator.GetGlobalValidator()

	if err := globalValidator.ValidateStructModel(r); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r DoTransferRequest) ToTransferInput(
	accOriginID string,
) transfers.TransferInput {
	return transfers.TransferInput{
		AccountOriginID:      accOriginID,
		AccountDestinationID: r.AccountDestinationID,
		Amount:               r.Amount,
	}
}

// DoTransfer godoc
//
//	@Summary		Realizar transferência entre contas.
//	@Description	Endpoint utilizado para realizar uma transferência entre contas.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			request		body		DoTransferRequest	true	"Dados da transação"
//	@Param			account_id	path		string				true	"Identificador da conta de origem"
//	@Success		200			{object}	TransferResponse	"Transferência realizada com sucesso."
//	@Failure		400			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse	"Conta não encontrada."
//	@Failure		422			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/api/v1/challenge-accounts/accounts/{account_id}/transfers [post]
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

	return response.Ok(TransferResponse{
		ID:                   out.ID,
		AccountOriginID:      out.AccountOriginID,
		AccountDestinationID: out.AccountDestinationID,
		Amount:               out.Amount,
		CreatedAt:            out.CreatedAt.Format(time.RFC3339),
	})
}
