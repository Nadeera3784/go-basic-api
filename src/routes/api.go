package routes

import (
	"application/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api/v1") 
	api.Get("/books", controllers.GetAll)
	api.Get("/books/:id", controllers.GetById)
	api.Post("/books", controllers.Create)
	api.Put("/books/:id", controllers.Update)
	api.Delete("/books/:id", controllers.Delete)
}