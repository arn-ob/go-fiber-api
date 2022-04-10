package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/param", func(c *fiber.Ctx) error {
		return c.SendString("param: ")
	})

	app.Get("/id/:id", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("id"))
	})

	app.Listen(":3080")
}
