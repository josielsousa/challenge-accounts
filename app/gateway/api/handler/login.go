package handler

import (
	"net/http"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
)

func (Handler) Login(_ *http.Request) *response.Response {
	return response.NoContent()
}
