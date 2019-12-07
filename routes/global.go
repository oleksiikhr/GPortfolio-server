package routes

import (
	"net/http"
)

// GlobalRoutes main routes for Frontend parts
func (app App) globalRoutes() {
	http.Handle("/static/", http.StripPrefix("/", http.FileServer(http.Dir("dist"))))
	http.HandleFunc("/", app.handleMain)
}

// handleMain handler for displaying the page
func (app App) handleMain(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/main.js", "/static/main.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(app.Html)
}
