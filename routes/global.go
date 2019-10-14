package routes

import (
	"encoding/json"
	"net/http"
)

// GlobalRoutes main routes for Frontend parts
func (app App) GlobalRoutes() {
	http.Handle("/static/", http.StripPrefix("/", http.FileServer(http.Dir("dist"))))
	http.HandleFunc("/", app.handleMain)
}

// handleMain handler for displaying the page
func (app App) handleMain(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/main.js", "/static/main.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(app.Html)
}

// push to speed up content delivery (http/2)
func push(w http.ResponseWriter, resources ...string) {
	if pusher, ok := w.(http.Pusher); ok {
		for _, resource := range resources {
			_ = pusher.Push(resource, nil)
		}
	}
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
