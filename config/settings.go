package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system variables - ", err.Error())
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// AccessJWTExpiryDuration 15min
var AccessJWTExpiryDuration = time.Hour * 15

// RefreshJWTExpiryDuration 7days
var RefreshJWTExpiryDuration = time.Hour * 24 * 7
