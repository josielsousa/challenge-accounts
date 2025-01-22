package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/validator"
)

type (
	CreateAccountRequest struct {
		Name     string `json:"name"     validate:"required"`
		Balance  int    `json:"balance"  validate:"required"`
		CPF      string `json:"cpf"      validate:"required,cpf"`
		Password string `json:"password" validate:"required"`
	}

	AccountResponse struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Balance   int    `json:"balance"`
		CPF       string `json:"cpf"`
		CreatedAt string `json:"created_at"`
	}
)

func (r CreateAccountRequest) Validate() error {
	globalValidator := validator.GetGlobalValidator()

	if err := globalValidator.ValidateStructModel(r); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r CreateAccountRequest) ToCreateAccountInput() accounts.AccountInput {
	return accounts.AccountInput{
		Name:    r.Name,
		Balance: r.Balance,
		CPF:     r.CPF,
		Secret:  r.Password,
	}
}

func (h Handler) CreateAccount(req *http.Request) *response.Response {
	var input CreateAccountRequest

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		return response.BadRequest(err)
	}

	if err := input.Validate(); err != nil {
		return response.BadRequest(err)
	}

	out, err := h.accUC.Create(req.Context(), input.ToCreateAccountInput())
	if err != nil {
		return response.AppError(err)
	}

	return response.Ok(AccountResponse{
		ID:        out.ID,
		Name:      out.Name,
		Balance:   out.Balance,
		CPF:       out.CPF,
		CreatedAt: out.CreatedAt.Format(time.RFC3339),
	})
}
