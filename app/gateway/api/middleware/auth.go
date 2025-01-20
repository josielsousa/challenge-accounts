package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
	"github.com/josielsousa/challenge-accounts/types"
)

func Authorize(signer *jwt.Jwt, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		claims, err := signer.Authorize(req.Header.Get("Authorization"))
		if err != nil {
			slog.Error("on authorize", slog.Any("error", err), slog.String("path", req.URL.Path))

			writer.WriteHeader(http.StatusUnauthorized)
			writer.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(writer).Encode(response.UnauthorizedErr)
			if err != nil {
				slog.Error("write response", slog.Any("error", err))

				return
			}

			return
		}

		ctx := types.ContextWithClaims(req.Context(), claims)

		next.ServeHTTP(writer, req.WithContext(ctx))
	})
}
