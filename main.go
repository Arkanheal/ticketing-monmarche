package main

import (
	"log"

    "tickets/router"
    "tickets/database"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"

    _ "github.com/lib/pq"
)

func main() {
    // Connexion à la base
    if err := database.Connect(); err != nil {
        log.Fatal(err)
    }

    // Création de l'app
    app := fiber.New()

    // Middleware de gestion d'erreur
    app.Use(recover.New(recover.Config{
        EnableStackTrace: true, // False en prod (config file?)
    }))

    // Middleware de log
    app.Use(logger.New(logger.Config{
        Format: "${latency} - [${ip}:${port}] ${status} - ${method} ${path}\n",
        // Output:
    }))

    router.SetupRoutes(app)

    // Lancement de l'app
    app.Listen("localhost:3000")
}
