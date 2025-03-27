package router

import (
	health "simple-oauth2-server/internal/api/handler"

	"github.com/go-chi/chi"
)


func New() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Get("/health", health.Get)

	return r, nil
}
