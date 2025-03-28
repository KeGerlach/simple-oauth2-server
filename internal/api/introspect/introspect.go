package introspect

import (
	"encoding/json"
	"net/http"
	"simple-oauth2-server/internal/environment"
	"strings"

	"maps"

	"github.com/golang-jwt/jwt/v4"
)

func Post(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusBadRequest)
			return
	}

	splittedAuthHeader := strings.Split(authHeader, " ")
	if len(splittedAuthHeader) != 2 || splittedAuthHeader[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusBadRequest)
			return
	}

	token := splittedAuthHeader[1]

	claims := jwt.MapClaims{}
	response := map[string]any{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return environment.Get().PUBLIC_KEY, nil
	})
	
	if err != nil || !parsedToken.Valid {
		response["active"] = false
	} else {
		response["active"] = true
		maps.Copy(response, claims)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
	}
}