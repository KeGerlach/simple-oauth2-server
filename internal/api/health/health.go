package health

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Status	string	`json:"status"`
}

func Get(w http.ResponseWriter, _ *http.Request) {
	body := Response{
		Status: "ok",
	}

	response, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
	}

	w.Write(response)
}