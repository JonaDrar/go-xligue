package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		//imprimir el log del error
		log.Fatal("Error loading .env file")
		fmt.Println(err)
	}

	return os.Getenv("MONGODB_URI")
}
