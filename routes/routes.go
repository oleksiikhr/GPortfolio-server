package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis/v7"
)

// App this is the core of the application
type App struct {
	Redis  *redis.Client
	Logger *log.Logger
	Html   []byte
}

// NewRoutes - register all exists routes
func (app App) NewRoutes() {
	app.globalRoutes()
	app.githubRoutes()
}

// push to speed up content delivery (http/2)
func push(w http.ResponseWriter, resources ...string) {
	if pusher, ok := w.(http.Pusher); ok {
		for _, resource := range resources {
			_ = pusher.Push(resource, nil)
		}
	}
}

// TODO
func response(w http.ResponseWriter, msg string, statusCode int) map[string]interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := make(map[string]interface{})
	response["message"] = msg

	return response
}

// respJsonFailed set badRequest to response and output message of error
func respJsonFailed(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := make(map[string]interface{})
	response["message"] = msg
	json.NewEncoder(w).Encode(response)
}

// respJsonSuccess with message
func respJsonSuccess(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := make(map[string]interface{})
	response["message"] = msg
	json.NewEncoder(w).Encode(response)
}
