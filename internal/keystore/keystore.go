package keystore

import "fmt"

type KeyStore struct {
	Keys map[string]KeyPair
}

func New() *KeyStore {
	ks := &KeyStore{Keys: make(map[string]KeyPair)}

	if err := ks.GenerateKeyPair(); err != nil {
		fmt.Printf("Something went wrong during key pair generation")
	}

	return ks
}