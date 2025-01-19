package rest

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/josielsousa/challenge-accounts/app/gateway/http/rest/response"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderApplicationJSON = "application/json"
)

type handlerFunc func(req *http.Request) *response.Response

func Handler(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := fn(r)

		if err := WriteJSON(w, resp.StatusCode, resp.Body); err != nil {
			slog.Error("write response", slog.Any("error", err))
		}
	}
}

func WriteJSON(rw http.ResponseWriter, statusCode int, payload any) error {
	rw.WriteHeader(statusCode)

	if payload == nil {
		return nil
	}

	rw.Header().Set(HeaderContentType, HeaderApplicationJSON)

	if err := json.NewEncoder(rw).Encode(payload); err != nil {
		return fmt.Errorf("encode response body: %w", err)
	}

	return nil
}
