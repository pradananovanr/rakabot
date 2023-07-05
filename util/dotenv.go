package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("failed to load .env files")
	}

	return os.Getenv(key)
}
