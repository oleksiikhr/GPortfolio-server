package routes

import (
	"net/http"
)

// globalRoutes main routes for Frontend parts
func (h *Handlers) globalRoutes() {
	http.Handle("/static/", http.StripPrefix("/", http.FileServer(http.Dir("dist"))))
	http.HandleFunc("/", h.handleMain())
}

// handleMain handler for displaying the page
func (h *Handlers) handleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		push(w, "/static/main.js", "/static/main.css")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(h.Html)
	}
}
