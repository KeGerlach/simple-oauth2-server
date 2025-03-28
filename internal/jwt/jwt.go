package jwt

import (
	"fmt"
	"simple-oauth2-server/internal/environment"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Generate(clientID string) (string, int, error) {
	expiresIn:= environment.Get().TOKEN_EXPIRATION_TIME
	expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)

	claims := jwt.RegisteredClaims{
		Subject: 	clientID,
		ExpiresAt: 	jwt.NewNumericDate(expiresAt),
		IssuedAt: 	jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(environment.Get().PRIVATE_KEY)
	if err != nil {
		fmt.Printf("%s", err)
		return "", 0, err
	}

	return signedToken, expiresIn, nil
}
