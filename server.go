package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"fiber-api/config"
    "fiber-api/handlers"
)

func main() {
	app := fiber.New()

	// gorm connect
	config.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

    app.Get("/dogs", handlers.GetDogs)
    app.Get("/dogs/:id", handlers.GetDog)
    app.Post("/dogs", handlers.AddDog)
    app.Put("/dogs/:id", handlers.UpdateDog)
    app.Delete("/dogs/:id", handlers.RemoveDog)

	log.Fatal(app.Listen(":3080"))
}
