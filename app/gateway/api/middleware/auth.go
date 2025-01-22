package middleware

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
	"github.com/josielsousa/challenge-accounts/types"
)

func Authorize(signer *jwt.Jwt, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		claims, err := signer.Authorize(req.Header.Get("Authorization"))
		if err != nil {
			slog.Error("on authorize", slog.Any("error", err), slog.String("path", req.URL.Path))

			resp := response.Unauthorized()

			render.Status(req, resp.StatusCode)
			render.JSON(rw, req, resp.Body)

			return
		}

		ctx := types.ContextWithClaims(req.Context(), claims)

		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
