package handlers

import (
    "github.com/gofiber/fiber/v3"
    "belajar-fiber/models"
)

func GetUsers(c fiber.Ctx) error {
    users := []models.User{
        {
            ID:    1,
            Name:  "Arif",
            Email: "arif@example.com",
        },
        {
            ID:    2,
            Name:  "Budi",
            Email: "budi@example.com",
        },
    }

    return c.JSON(users)
}