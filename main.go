package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jonadrar/go-xligue/configs"
	"github.com/jonadrar/go-xligue/routes"
)

func main() {
	app := fiber.New()

	//run DB connection
	configs.ConnectDB()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	//routes
	routes.UserRoute(app)

	app.Listen(":" + port)
}
