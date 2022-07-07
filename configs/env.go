package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoUrl() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGO_URL")
}
