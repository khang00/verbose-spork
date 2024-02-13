package handler

import (
	"encoding/json"
	"net/http"
)

func AllowCORS(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			json.NewEncoder(w).Encode("OKOK")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			endpointHandler(w, r)
		}
	}
}
