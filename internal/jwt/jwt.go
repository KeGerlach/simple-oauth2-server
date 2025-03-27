package jwt

import (
	"fmt"
	"simple-oauth2-server/internal/environment"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Generate(clientID string) (string, int, error) {
	expiresIn:= environment.Get().TOKEN_EXPIRATION_TIME
	expiresAt := time.Now().Add(time.Duration(expiresIn))

	claims := jwt.RegisteredClaims{
		Subject: 	clientID,
		ExpiresAt: 	jwt.NewNumericDate(expiresAt),
		IssuedAt: 	jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(environment.Get().SECRET)
	if err != nil {
		fmt.Printf("%s", err)
		return "", 0, err
	}

	return signedToken, expiresIn, nil
}

// func loadPrivateKey() (*rsa.PrivateKey, error) {
// 	bytes, err := os.ReadFile(environment.Get().PRIVATE_KEY_PATH)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read private key: %w", err)
// 	}

// 	block, _ := pem.Decode(bytes)
// 	if block == nil {
// 		return nil, fmt.Errorf("no PEM data found")
// 	}

// 	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse private key: %w", err)
// 	}

// 	return key, nil
// }
