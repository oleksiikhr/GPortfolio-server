package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GPortfolio/server/config"
	"github.com/GPortfolio/server/routes"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	// Connect to Redis
	rClient, err := cRedis()
	if err != nil {
		log.Fatal(err)
	}

	// Load html file
	html, err := loadHomePageHtml()
	if err != nil {
		log.Fatal(err)
	}

	// Run application with routes
	app := routes.App{Redis: rClient, Html: html}
	app.GlobalRoutes()
	app.GithubRoutes()

	// Run server
	startServer()
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

// loadHomePageHtml it's index.html from dist folder (frontend/web after build)
func loadHomePageHtml() ([]byte, error) {
	return ioutil.ReadFile("./dist/index.html")
}

// startServer starts the server on http/https depending on the environment
func startServer() {
	addr := config.Env("APP_ADDR", "localhost:8080")
	fmt.Println("Server Run:", addr)
	var err error

	if config.Env("APP_TLS", "false") == "true" {
		err = http.ListenAndServeTLS(addr, config.Env("APP_TLS_CERT", ""), config.Env("APP_TLS_KEY", ""), nil)
	} else {
		err = http.ListenAndServe(addr, nil)
	}

	if err == nil {
		fmt.Println("Server stopped")
	} else {
		fmt.Println(err)
	}
}
