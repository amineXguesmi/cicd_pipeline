package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongoDBURI string
	DBName     string
	JWTSecret  string
)

func LoadEnv() {
	// Skip loading .env if the environment is set to "test"
	if os.Getenv("ENV") != "test" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file:", err)
		}
	}
	
	MongoDBURI = os.Getenv("MONGODB_URI")
	JWTSecret = os.Getenv("JWT_SECRET")
	DBName = os.Getenv("DB_NAME")
}
