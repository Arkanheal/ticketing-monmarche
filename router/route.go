package router

import (
    "github.com/gofiber/fiber/v2"
    "tickets/handler"
)

func SetupRoutes(app *fiber.App) {

    // Un seul endpoint donc c'est facile
    app.Post("/ticket", handler.CreateOrder)
}
