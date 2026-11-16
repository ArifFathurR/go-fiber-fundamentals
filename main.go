package main

import (
	"belajar-fiber/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
    // Membuat instance aplikasi Fiber
    app := fiber.New()

    // Membuat endpoint GET
    // app.Get("/", func(c fiber.Ctx) error {
    //     return c.SendString("Hello, Welcome to Go Fiber!")
    // })

    routes.SetupRoutes(app)

    // Menjalankan server
    log.Fatal(app.Listen(":3000"))
}