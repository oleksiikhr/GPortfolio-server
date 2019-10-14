package api

import (
	"encoding/json"
	"net/http"
)

// respJsonFailed set badRequest to response and output message of error
func respJsonFailed(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := make(map[string]interface{})
	response["message"] = msg
	json.NewEncoder(w).Encode(response)
}

// responseSuccess with message
func jsonRespSuccess(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := make(map[string]interface{})
	response["message"] = msg
	json.NewEncoder(w).Encode(response)
}
