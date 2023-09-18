package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonadrar/go-xligue/controllers"
)

func UserRoute(app *fiber.App) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Post("user", controllers.CreateUser)
	app.Get("user/:userId", controllers.GetAUser)
	app.Put("user/:userId", controllers.EditUser)
	app.Delete("user/:userId", controllers.DeleteAUser)
	app.Get("users", controllers.GetAllUsers)
}
