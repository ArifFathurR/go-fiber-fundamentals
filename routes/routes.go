package routes

import (
	"belajar-fiber/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App)  {
	app.Get("/users", handlers.GetUsers)
}