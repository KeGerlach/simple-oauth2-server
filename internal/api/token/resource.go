package token

import (
	"simple-oauth2-server/internal/environment"
	"simple-oauth2-server/internal/keystore"
)

type Resource struct {
	env *environment.Environment
	ks *keystore.KeyStore
}

func New(environment *environment.Environment, keystore *keystore.KeyStore) *Resource {
	return &Resource{
		env: environment,
		ks: keystore,
	}
}