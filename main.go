package main

import (
	"log"

    "tickets/router"
    "tickets/database"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/proxy"
    "github.com/gofiber/fiber/v2/middleware/recover"

    _ "github.com/lib/pq"
)

func main() {
    if err := database.Connect(); err != nil {
        log.Fatal(err)
    }

    app := fiber.New()

    app.Use(recover.New(recover.Config{
        EnableStackTrace: true,
    }))

    app.Use(logger.New(logger.Config{
        Format: "${latency} - [${ip}:${port}] ${status} - ${method} ${path}\n",
    }))

    router.SetupRoutes(app)

    app.Listen("localhost:3000")
}
