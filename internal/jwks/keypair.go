package jwks

import (
	"crypto/rand"
	"crypto/rsa"
)

type KeyPair struct {
	PrivateKey 	*rsa.PrivateKey
	PublicKey 	*rsa.PublicKey
	Kid			string
}

func (ks *KeyStore) GetActiveKeyPair() (KeyPair, error) {
	var keyPair KeyPair

	return keyPair, nil
}


func (ks *KeyStore) GenerateKeyPair() error {

	return nil
}

func GenerateRsa256Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}