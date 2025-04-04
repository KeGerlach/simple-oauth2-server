package jwks

type KeyStore struct {
	Keys map[string]KeyPair
}

// var keystore *KeyStore