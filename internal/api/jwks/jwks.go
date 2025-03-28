package jwks

import (
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"simple-oauth2-server/internal/environment"
)

type Jwk struct {
	Kty		string	`json:"kty"`
	N		string	`json:"n"`
	E		string	`json:"e"`
	Alg		string	`json:"alg"`
	Use		string	`json:"use"`
}

type Jwks struct {
	Keys	[]Jwk	`json:"keys"`	
}

func Get(w http.ResponseWriter, r *http.Request) {
	publicKey := environment.Get().PUBLIC_KEY

	keys := Jwks {
		Keys: []Jwk {
			{
				Kty:	"RSA",
				N: 		base64.RawURLEncoding.EncodeToString(publicKey.N.Bytes()),
				E:		base64.RawURLEncoding.EncodeToString(big.NewInt(int64(publicKey.E)).Bytes()),
				Alg: 	"RS256",
				Use: 	"sig",
			},
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(keys)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
	}
}