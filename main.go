package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GPortfolio/server/config"
	"github.com/GPortfolio/server/routes"
	"github.com/GPortfolio/server/services/redis"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	// Create a new logger
	logger := log.New(os.Stdout, config.ProjectName+" ", log.LstdFlags|log.Lshortfile)

	// Connect to Redis
	redisClient, err := redis.NewRedis(config.Env("REDIS_ADDR", config.DefaultRedisAddr))
	if err != nil {
		logger.Fatal(err)
	}

	// Run application with routes
	h := routes.Handlers{
		Redis:  redisClient,
		Logger: logger,
	}

	h.NewRoutes()

	// Run server
	startServer(logger)
}

// startServer starts the server on http/https depending on the environment
func startServer(logger *log.Logger) {
	var err error

	addr := config.Env("APP_ADDR", config.DefaultAppAddr)
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
