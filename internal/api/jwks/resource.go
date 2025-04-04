package jwks

import "simple-oauth2-server/internal/keystore"

type Resource struct {
	keystore *keystore.KeyStore;
}

func New(keystore *keystore.KeyStore) *Resource {
	return &Resource{
		keystore: keystore,
	}
}
