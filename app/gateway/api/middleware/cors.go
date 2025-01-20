package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

var allowedHeaders = []string{
	"Accept",
	"Authorization",
	"Content-Type",
	"Origin",
	"Referer",
	"User-Agent",
}

var allowedMethods = []string{
	"GET",
	"POST",
	"PATCH",
	"PUT",
	"DELETE",
	"OPTIONS",
	"HEAD",
}

func CORS() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   allowedMethods,
		AllowedHeaders:   allowedHeaders,
		AllowCredentials: false,
	})
}
