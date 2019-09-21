package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var htmlFile = []byte("")

func main() {
	// Load .env file
	_ = godotenv.Load()

	loadHtmlFiles()

	// Routes
	http.Handle("/static/", http.StripPrefix("/", http.FileServer(http.Dir("dist"))))
	http.HandleFunc("/", handleMain)

	startServer()
}

func loadHtmlFiles() {
	var err error
	htmlFile, err = ioutil.ReadFile("./dist/index.html")

	if err != nil {
		log.Fatal(err)
	}
}

func push(w http.ResponseWriter, resources ...string) {
	if pusher, ok := w.(http.Pusher); ok {
		for _, resource := range resources {
			_ = pusher.Push(resource, nil)
		}
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/main.js", "/static/main.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlFile)
}

func startServer() {
	addr := env("APP_ADDR", ":8080")

	if env("APP_TLS", "false") == "true" {
		http.ListenAndServeTLS(addr, env("APP_TLS_CERT", ""), env("APP_TLS_KEY", ""), nil)
	} else {
		http.ListenAndServe(addr, nil)
	}
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
