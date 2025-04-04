package introspect

import "simple-oauth2-server/internal/keystore"

type Resource struct {
	ks *keystore.KeyStore
}

func New(keystore *keystore.KeyStore) *Resource {
	return &Resource{
		ks: keystore,
	}
}
