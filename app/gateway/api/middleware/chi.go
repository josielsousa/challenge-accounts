package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest/response"
)

var (
	StripSlashes = middleware.StripSlashes
	CleanPath    = middleware.CleanPath
)

// Recoverer - middleware to recover from panic.
// based on default middleware from chi, with custom internal server error
// response.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer func(ctx context.Context) {
			if rec := recover(); rec != nil {
				err := fmt.Errorf("%v", rec)

				resp := response.InternalServerErr(err)

				render.Status(req, resp.StatusCode)
				render.JSON(rw, req, resp.Body)

				slog.ErrorContext(
					ctx,
					fmt.Sprintf("recover middleware -> %v", err.Error()),
				)
			}
		}(context.WithoutCancel(req.Context()))

		next.ServeHTTP(rw, req)
	})
}
