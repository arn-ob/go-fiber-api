package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"fiber-api/config"
	"fiber-api/routes"
)

func main() {
	app := fiber.New()

	// gorm connect
	config.Connect()

	// config routes
	routes.Routes(app)

	log.Fatal(app.Listen(":3080"))
}
