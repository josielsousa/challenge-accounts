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
		// Name - nome do titular da conta.
		Name string `json:"name" validate:"required"`
		// Balance - saldo inicial da conta.
		Balance int `json:"balance" validate:"required"`
		// CPF - número do CPF do titular da conta.
		CPF string `json:"cpf" validate:"required,cpf"`
		// Password - senha de acesso a conta.
		Password string `json:"password" validate:"required"`
	}

	AccountResponse struct {
		// ID - identificador único da conta.
		ID string `json:"id"`
		// Name - nome do titular da conta.
		Name string `json:"name"`
		// Balance - saldo atual da conta.
		Balance int `json:"balance"`
		// CPF - número do CPF do titular da conta.
		CPF string `json:"cpf"`
		// CreatedAt - data de criação da conta.
		CreatedAt string `json:"created_at"`
	} //	@name	CreateAccountResponse
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

// CreateAccount godoc
//
//	@Summary		Criar nova conta.
//	@Description	Endpoint utilizado para criação de uma nova conta.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateAccountRequest	true	"Dados da conta"
//	@Success		201		{object}	CreateAccountResponse	"Conta criada com sucesso."
//	@Failure		400		{object}	ErrorResponse
//	@Failure		422		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/v1/challenge-accounts/accounts [post]
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

	return response.Created(AccountResponse{
		ID:        out.ID,
		Name:      out.Name,
		Balance:   out.Balance,
		CPF:       out.CPF,
		CreatedAt: out.CreatedAt.Format(time.RFC3339),
	})
}
