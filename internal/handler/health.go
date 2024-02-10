package handler

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	response := HealthResponse{Message: "ok"}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
