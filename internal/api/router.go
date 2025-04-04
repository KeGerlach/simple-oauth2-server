package router

import (
	"simple-oauth2-server/internal/api/health"
	"simple-oauth2-server/internal/api/introspect"
	"simple-oauth2-server/internal/api/jwks"
	"simple-oauth2-server/internal/api/token"
	"simple-oauth2-server/internal/keystore"

	"github.com/go-chi/chi"
)


func New(ks *keystore.KeyStore) (*chi.Mux, error) {
	r := chi.NewRouter()
	
	// init handlers
	jwks := jwks.New(ks)

	// routes
	r.Get("/health", health.Get)

	r.Post("/token", token.Post)

	r.Get("/jwks", jwks.Get)

	r.Post("/introspect", introspect.Post)

	return r, nil
}
