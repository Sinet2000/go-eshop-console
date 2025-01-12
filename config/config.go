package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func GetEnv(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Fatalf("Environment variable %s not set", key)
    }
    return value
}
