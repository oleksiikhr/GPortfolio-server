package routes

import (
	"log"
	"math/rand"
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

func (h *Handlers) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerSecurityPass := r.Header.Get("Security-Pass")
		headerSecurityKey := r.Header.Get("Security-Key")

		if headerSecurityKey != "" && headerSecurityPass != "" {
			data := h.Redis.Get(headerSecurityKey)
			h.Logger.Println(data)
			// if exists
			// 	set app.User = redis.data
		}

		w.Header().Set("Security-Key", rnd(25))
		w.Header().Set("Security-Pass", rnd(60))

		h.Logger.Println("HERE", r.URL)
		next.ServeHTTP(w, r)
	}
}

// TODO Move to other package
// rnd generate a random of symbols specified length
func rnd(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letterRunesLen := len(letterRunes)
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(letterRunesLen)]
	}

	return string(b)
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
