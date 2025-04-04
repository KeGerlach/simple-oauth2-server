package keystore

import (
	"encoding/base64"
	"math/big"
)

type Jwk struct {
	Kty		string	`json:"kty"`
	N		string	`json:"n"`
	E		string	`json:"e"`
	Alg		string	`json:"alg"`
	Use		string	`json:"use"`
	Kid		string 	`json:"kid"`
}

type Jwks struct {
	Keys	[]Jwk	`json:"keys"`	
}

func (ks *KeyStore) ToJwks() Jwks {
	jwks := Jwks{Keys: []Jwk{}}

	for _, key := range ks.Keys {
		jwks.Keys = append(jwks.Keys, Jwk{
				Kty:	"RSA",
				N: 		base64.RawURLEncoding.EncodeToString(key.PublicKey.N.Bytes()),
				E:		base64.RawURLEncoding.EncodeToString(big.NewInt(int64(key.PublicKey.E)).Bytes()),
				Alg: 	"RS256",
				Use: 	"sig",
				Kid: 	key.Kid,
			})
	}

	return jwks
}