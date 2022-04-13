package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Ok(c *fiber.Ctx) error {
    return c.Status(200).SendString("Respose Ok")
}