package handler

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if data == `{}` {
		w.Write([]byte(`{}`))
	} else {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.Write([]byte(data.(string)))
		}
	}
}
