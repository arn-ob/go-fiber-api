package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber-api/handlers"
)


func Routes (app *fiber.App) {
	
	// Testing Routing
	app.Get("/", handlers.Ok)

	// crud 
	app.Get("/dogs", handlers.GetDogs)
    app.Get("/dogs/:id", handlers.GetDog)
    app.Post("/dogs", handlers.AddDog)
    app.Put("/dogs/:id", handlers.UpdateDog)
    app.Delete("/dogs/:id", handlers.RemoveDog)
}