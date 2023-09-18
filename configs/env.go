package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {

	if os.Getenv("ENVIRONMENT") != "production" {
		// Solo carga el archivo .env en entornos no productivos
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	return os.Getenv("MONGODB_URI")
}
