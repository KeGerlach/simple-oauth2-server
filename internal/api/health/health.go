package health

import "net/http"

func Get(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(`{"status": "ok"}`))
}