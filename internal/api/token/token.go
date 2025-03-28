package token

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"simple-oauth2-server/internal/environment"
	"simple-oauth2-server/internal/jwt"
	"strings"
)

type Response struct{
	AccessToken	string 	`json:"access_token"`
	TokenType	string 	`json:"token_type"`
	ExpiresIn	int		`json:"expires_in"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}

	splittedAuthHeader := strings.Split(authHeader, " ")
	if len(splittedAuthHeader) != 2 || splittedAuthHeader[0] != "Basic" {
		http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(splittedAuthHeader[1])
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 {
		http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
	}

	clientId, clientSecret := credentials[0], credentials[1]

	if clientId != environment.Get().CLIENT_ID || clientSecret != environment.Get().CLIENT_SECRET {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, expiresIn, err := jwt.Generate(clientId)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusUnauthorized)
		return
	}

	response := Response{
		AccessToken: 	token,
		TokenType: 		"Bearer",
		ExpiresIn: 		expiresIn,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
	}
}