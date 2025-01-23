package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/validator"
	"github.com/josielsousa/challenge-accounts/types"
)

type (
	SinginRequest struct {
		Cpf      string `json:"cpf"      validate:"required,cpf"`
		Password string `json:"password" validate:"required"`
	}

	SinginReponse struct {
		Token string `json:"token"`
	}
)

func (r SinginRequest) Validate() error {
	globalValidator := validator.GetGlobalValidator()

	if err := globalValidator.ValidateStructModel(r); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (h Handler) Signin(req *http.Request) *response.Response {
	var request SinginRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return response.BadRequest(err)
	}
	defer req.Body.Close()

	if err := request.Validate(); err != nil {
		return response.BadRequest(err)
	}

	auth, err := h.authUC.Signin(req.Context(), types.Credentials{
		Cpf:    request.Cpf,
		Secret: request.Password,
	})
	if err != nil {
		return response.AppError(err)
	}

	return response.Ok(SinginReponse{Token: auth.Token})
}
