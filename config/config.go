package config

import "os"

const (
	// ProjectName from Github repository
	ProjectName = "GPortfolio"
	// HomePageFile path dist from web
	HomePageFile = "./dist/index.html"
	// DefaultRedisAddr connect to the redis by this address if not present
	DefaultRedisAddr = "localhost:6379"
	// DefaultAppAddr run application by this address if not present
	DefaultAppAddr = "localhost:8080"
)

// Env get environment of system, can be modified in .env file
func Env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// GetUrlWebsite of this project
func GetUrlWebsite() string {
	var website string

	if Env("APP_TLS", "false") == "true" {
		website = "https://"
	} else {
		website = "http://"
	}

	return website + Env("APP_ADDR", "")
}
