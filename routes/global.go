package routes

import (
	"net/http"
)

// globalRoutes main routes for Frontend parts
func (h *Handlers) globalRoutes() {
	http.Handle("/", http.FileServer(http.Dir("dist")))
}
