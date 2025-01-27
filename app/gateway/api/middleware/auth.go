package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/render"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
	"github.com/josielsousa/challenge-accounts/app/gateway/jwt"
	"github.com/josielsousa/challenge-accounts/app/types"
)

const minBearerTokenParts = 2

func Authorize(signer *jwt.Jwt, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		logger := slog.Default().With(slog.String("path", req.URL.Path))

		token, ok := sanitizeBearerToken(req.Header.Get("Authorization"))
		if !ok {
			logger.Error(
				"on sanitize token",
				slog.String("token", token),
				slog.Bool("ok", ok),
			)

			handleUnauthorized(req, rw)

			return
		}

		claims, err := signer.Authorize(token)
		if err != nil {
			logger.Error("on authorize", slog.Any("error", err))

			handleUnauthorized(req, rw)

			return
		}

		ctx := types.ContextWithClaims(req.Context(), claims)

		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func sanitizeBearerToken(authorization string) (string, bool) {
	if strings.HasPrefix(authorization, "Bearer") {
		bearerToken := strings.Split(authorization, " ")

		if len(bearerToken) != minBearerTokenParts {
			return "", false
		}

		return bearerToken[1], true
	}

	if len(authorization) == 0 {
		return "", false
	}

	return authorization, true
}

func handleUnauthorized(req *http.Request, rw http.ResponseWriter) {
	resp := response.Unauthorized()

	render.Status(req, resp.StatusCode)
	render.JSON(rw, req, resp.Body)
}
