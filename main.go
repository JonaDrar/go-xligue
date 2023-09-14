package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Static("/", "./public")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
			"message": "OK",
			"data":    "All users",
		})
	})

	app.Listen(":" + port)
	fmt.Println("Server is running on port 3000")
}
