package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/arn-ob/go-fiber-api/config"
    "github.com/arn-ob/go-fiber-api/handlers"
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

	app.Listen(":3080")
}
