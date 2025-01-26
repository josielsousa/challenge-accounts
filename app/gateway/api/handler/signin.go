package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/josielsousa/challenge-accounts/app/domain/auth"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/validator"
)

type (
	// @description Dados de identificação para o signin
	SinginRequest struct {
		// Cpf - número do CPF do titular da conta, usado como username para o signin.
		Cpf string `json:"cpf" validate:"required,cpf"`
		// Password - senha de acesso a conta.
		Password string `json:"password" validate:"required"`
	} // @name SinginRequest

	SinginReponse struct {
		// Token - token de autenticação usado para acessar endpoints privados.
		Token string `json:"token"`
	} // @name SinginReponse
)

func (r SinginRequest) Validate() error {
	globalValidator := validator.GetGlobalValidator()

	if err := globalValidator.ValidateStructModel(r); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// CreateAccount godoc
//
//	@Summary		Autoriza o acesso a conta.
//	@Description	Endpoint utilizado para autorizar o acesso a conta.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SinginRequest	true "Dados de identificação"
//	@Success		200		{object}	SinginReponse	"Conta criada com sucesso."
//	@Failure		400		{object}	ErrorResponse
//	@Failure		422		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/v1/challenge-accounts/auth/signin [post]
func (h Handler) Signin(req *http.Request) *response.Response {
	var request SinginRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return response.BadRequest(err)
	}
	defer req.Body.Close()

	if err := request.Validate(); err != nil {
		return response.BadRequest(err)
	}

	out, err := h.authUC.Signin(req.Context(), auth.SiginInput{
		Cpf:    request.Cpf,
		Secret: request.Password,
	})
	if err != nil {
		return response.AppError(err)
	}

	return response.Ok(SinginReponse{Token: out.Token})
}
