package middleware

import "github.com/go-chi/chi/v5/middleware"

var (
	StripSlashes = middleware.StripSlashes
	Recoverer    = middleware.Recoverer
	CleanPath    = middleware.CleanPath
)
