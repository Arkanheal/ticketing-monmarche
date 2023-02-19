package router

import (
    "github.com/gofiber/fiber/v2"
    "tickets/handler"
)

func SetupRoutes(app *fiber.App) {

    app.Post("/ticket", handler.CreateOrder)
}
