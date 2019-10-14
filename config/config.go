package config

import "os"

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
