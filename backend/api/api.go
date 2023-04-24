package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func sendJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response := Response{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
		Data:    payload,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func sendError(w http.ResponseWriter, statusCode int, errorMessage string) {
	response := Response{
		Status:  statusCode,
		Message: errorMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
