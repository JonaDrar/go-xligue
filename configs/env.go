package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {

	//saber todos los archivos que tengo en el path
	fmt.Println(os.Environ())

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return os.Getenv("MONGODB_URI")
}
