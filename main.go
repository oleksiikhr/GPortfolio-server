package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/GPortfolio/server/config"
	"github.com/GPortfolio/server/routes"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

const (
	projectName      = "GPortfolio"
	homePageFile     = "./dist/index.html"
	defaultRedisAddr = "localhost:6379"
	defaultAppAddr   = "localhost:8080"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	// Create a new logger
	logger := log.New(os.Stdout, projectName+" ", log.LstdFlags|log.Lshortfile)

	// Connect to Redis
	redisClient, err := newRedis()
	if err != nil {
		logger.Fatal(err)
	}

	// Load html file
	html, err := loadHomePageHtml()
	if err != nil {
		logger.Fatal(err)
	}

	// Run application with routes
	app := routes.App{
		Redis:  redisClient,
		Logger: logger,
		Html:   html,
	}

	app.NewRoutes()

	// Run server
	startServer(logger)
}

// newRedis create a new connection to Redis client
func newRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.Env("REDIS_ADDR", defaultRedisAddr),
	})

	// Check connection
	_, err := client.Ping().Result()

	return client, err
}

// loadHomePageHtml it's index.html from dist folder (frontend/web after build)
func loadHomePageHtml() ([]byte, error) {
	return ioutil.ReadFile(homePageFile)
}

// startServer starts the server on http/https depending on the environment
func startServer(logger *log.Logger) {
	var err error

	addr := config.Env("APP_ADDR", defaultAppAddr)
	logger.Println("Server Running (" + addr + ")")

	if config.Env("APP_TLS", "false") == "true" {
		err = http.ListenAndServeTLS(addr, config.Env("APP_TLS_CERT", ""), config.Env("APP_TLS_KEY", ""), nil)
	} else {
		err = http.ListenAndServe(addr, nil)
	}

	if err == nil {
		logger.Println("Server Stopped")
	} else {
		logger.Println(err)
	}
}
