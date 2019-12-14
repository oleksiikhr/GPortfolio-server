package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GPortfolio/server/services/redis"
)

// Handlers
type Handlers struct {
	Redis  *redis.Redis
	Logger *log.Logger
}

// NewRoutes register all exists routes
func (h *Handlers) NewRoutes() {
	h.globalRoutes()
	h.githubRoutes()
}

// tryAccessToken from response headers
func (h *Handlers) tryAccessToken(r *http.Request) string {
	pass := r.Header.Get("Security-Pass")
	key := r.Header.Get("Security-Key")
	if key == "" || pass == "" {
		return ""
	}

	value := h.Redis.SecGetHard(key, pass)
	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}

// response json by specific structure
func response(w http.ResponseWriter, msg interface{}, statusCode int) map[string]interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := make(map[string]interface{})
	response["message"] = msg
	response["data"] = nil

	return response
}

// responseQuick json by specific structure
func responseQuick(w http.ResponseWriter, msg interface{}, statusCode int) {
	response := response(w, msg, statusCode)
	json.NewEncoder(w).Encode(response)
}
