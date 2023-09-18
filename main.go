package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonadrar/go-xligue/configs"
	"github.com/jonadrar/go-xligue/routes"
)

func main() {
	app := fiber.New()

	//run DB connection
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen(":3000")
}
