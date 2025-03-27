package jwt

import (
	"simple-oauth2-server/internal/environment"
	"strconv"
)

func Generate(clientID string) (string, int, error) {
	expiresIn, err := strconv.Atoi(environment.Get().TOKEN_EXPIRATION_TIME)
	if err != nil {
		expiresIn = 300
	}
	
	token := "tmp"

	return token, expiresIn, nil
} 