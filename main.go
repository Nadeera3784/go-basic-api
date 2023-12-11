package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "application/src/routes"
    "application/src/config"
)


func main() {

    app := fiber.New();

    app.Use(logger.New());

	config.SetupDatabase()

    routes.Setup(app);

    log.Println("API is running!");

    app.Listen(":4000")
}