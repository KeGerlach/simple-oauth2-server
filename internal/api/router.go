package router

import (
	health "simple-oauth2-server/internal/api/health"
	"simple-oauth2-server/internal/api/jwks"
	"simple-oauth2-server/internal/api/token"

	"github.com/go-chi/chi"
)


func New() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Get("/health", health.Get)

	r.Post("/token", token.Post)

	r.Get("/jwks", jwks.Get)

	return r, nil
}
