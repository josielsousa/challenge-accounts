package rest

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
)

type handlerFunc func(req *http.Request) *response.Response

func Handler(fn handlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		resp := fn(req)

		render.Status(req, resp.StatusCode)
		render.JSON(rw, req, resp.Body)
	}
}
