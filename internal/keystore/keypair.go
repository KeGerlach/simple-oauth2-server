package keystore

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/google/uuid"
)

type KeyPair struct {
	PrivateKey 	*rsa.PrivateKey
	PublicKey 	*rsa.PublicKey
	Kid			string
}

func (ks *KeyStore) GetActiveKeyPair() (KeyPair, error) {
	keyPair := KeyPair{}

	if len(ks.Keys) == 0 {
		return keyPair, fmt.Errorf("no keys available")
	}

	// for now only return last key pair found as active key pair
	for _, key := range ks.Keys {
		keyPair = key
	}

	return keyPair, nil
}

func (ks *KeyStore) GenerateKeyPair() error {
	privateKey, err := GenerateRsa256Key()
	if err != nil {
		return err
	}

	kid := uuid.New().String()
	ks.Keys[kid] = KeyPair{
		PrivateKey: privateKey,
		PublicKey: &privateKey.PublicKey,
		Kid: kid,
	}

	return nil
}

func GenerateRsa256Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}