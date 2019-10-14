package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/GPortfolio/server/api"
	"github.com/GPortfolio/server/config"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	// Connect to Redis
	redis, err := cRedis()
	if err != nil {
		log.Fatal(err)
	}

	// Load html file
	html, err := loadHomePageHtml()
	if err != nil {
		log.Fatal(err)
	}

	// Frontend Routes
	http.Handle("/static/", http.StripPrefix("/", http.FileServer(http.Dir("dist"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleMain(w, html)
	})

	// Api Routes
	api.GithubRoutes(redis)

	// Run server
	startServer()
}

// handleMain handler for displaying the page
func handleMain(w http.ResponseWriter, html []byte) {
	push(w, "/static/main.js", "/static/main.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(html)
}

// loadHomePageHtml it's index.html from dist folder (frontend/web after build)
func loadHomePageHtml() ([]byte, error) {
	return ioutil.ReadFile("./dist/index.html")
}

// newRedis create a new connection to Redis client
func cRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.Env("REDIS_ADDR", "localhost:6379"),
	})

	// Check connection
	_, err := client.Ping().Result()

	return client, err
}

// push to speed up content delivery (http/2)
func push(w http.ResponseWriter, resources ...string) {
	if pusher, ok := w.(http.Pusher); ok {
		for _, resource := range resources {
			_ = pusher.Push(resource, nil)
		}
	}
}

// startServer starts the server on http/https depending on the environment
func startServer() {
	addr := config.Env("APP_ADDR", "localhost:8080")
	fmt.Println("Server Run:", addr)

	if config.Env("APP_TLS", "false") == "true" {
		http.ListenAndServeTLS(addr, config.Env("APP_TLS_CERT", ""), config.Env("APP_TLS_KEY", ""), nil)
	} else {
		http.ListenAndServe(addr, nil)
	}
}

// rnd generate a random of symbols specified length
// https://stackoverflow.com/a/31832326/9612245
// TODO Is needed?
func rnd(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
