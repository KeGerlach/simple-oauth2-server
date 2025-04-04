package jwt

import (
	"fmt"
	"simple-oauth2-server/internal/environment"
	"simple-oauth2-server/internal/keystore"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Generate(env *environment.Environment, ks *keystore.KeyStore) (string, int, error) {
	expiresIn:= env.TOKEN_EXPIRATION_TIME
	expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)

	claims := jwt.RegisteredClaims{
		Subject: 	env.CLIENT_ID,
		ExpiresAt: 	jwt.NewNumericDate(expiresAt),
		IssuedAt: 	jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	keypair, err := ks.GetActiveKeyPair()
	if err != nil {
		fmt.Printf("%s", err)
		return "", 0, err
	}

	signedToken, err := token.SignedString(keypair.PrivateKey)
	if err != nil {
		fmt.Printf("%s", err)
		return "", 0, err
	}

	return signedToken, expiresIn, nil
}
