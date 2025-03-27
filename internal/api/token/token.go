package token

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-oauth2-server/internal/environment"
	"simple-oauth2-server/internal/jwt"
	"strings"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	token, err := jwt.Generate(clientId)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusUnauthorized)
		return
	}

	response := fmt.Appendf(nil, `{"accessToken": %s, "TokenType": "Bearer", "ExpiresIn": 42}`, token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}