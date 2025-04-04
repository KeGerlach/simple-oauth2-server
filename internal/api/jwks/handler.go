package jwks

import (
	"encoding/json"
	"net/http"
)

func (res *Resource) Get(w http.ResponseWriter, r *http.Request) {
	jwks := res.ks.ToJwks()
	
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jwks)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
	}
}