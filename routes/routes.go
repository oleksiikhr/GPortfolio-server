package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis/v7"
)

// Handlers
type Handlers struct {
	Redis  *redis.Client
	Logger *log.Logger
}

// NewRoutes register all exists routes
func (h *Handlers) NewRoutes() {
	h.globalRoutes()
	h.githubRoutes()
}

// push to speed up content delivery (http/2)
func push(w http.ResponseWriter, resources ...string) {
	if pusher, ok := w.(http.Pusher); ok {
		for _, resource := range resources {
			_ = pusher.Push(resource, nil)
		}
	}
}

// response json by specific structure
func response(w http.ResponseWriter, msg string, statusCode int) map[string]interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := make(map[string]interface{})
	response["message"] = msg
	response["data"] = nil

	return response
}

// responseQuick json by specific structure
func responseQuick(w http.ResponseWriter, msg string, statusCode int) {
	response := response(w, msg, statusCode)
	json.NewEncoder(w).Encode(response)
}
