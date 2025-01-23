package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
)

type handlerFunc func(req *http.Request) *response.Response

func Handler(fn handlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		resp := fn(req)

		err := sendJSON(rw, resp.Body, resp.StatusCode)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

func sendJSON(writer http.ResponseWriter, payload any, statusCode int) error {
	if payload == nil {
		writer.WriteHeader(statusCode)

		return nil
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	enc := json.NewEncoder(writer)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(payload); err != nil {
		return fmt.Errorf("encoding response: %w", err)
	}

	return nil
}
