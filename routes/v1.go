package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber-api/handlers"
)


func Routes (app *fiber.App) {
	
	// Testing Routing
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Testing Done")
	})

	app.Get("/dogs", handlers.GetDogs)
    app.Get("/dogs/:id", handlers.GetDog)
    app.Post("/dogs", handlers.AddDog)
    app.Put("/dogs/:id", handlers.UpdateDog)
    app.Delete("/dogs/:id", handlers.RemoveDog)
}