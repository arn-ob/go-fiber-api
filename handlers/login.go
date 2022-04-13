package handlers

import (
    "fiber-api/config"
    "fiber-api/entities"
    "github.com/gofiber/fiber/v2"
)

// Login registration list
func Registration(c *fiber.Ctx) error {
	
	registration := new(entities.Login)

    if err := c.BodyParser(registration); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    config.Database.Create(&registration)
    
	return c.Status(201).JSON(registration)
}

// Login user list
func Logins(c *fiber.Ctx) error {
	
	var login []entities.Login

	config.Database.Find(&login)

	return c.Status(200).JSON(login)
}

// Get Login Details
func GetLogin(c *fiber.Ctx) error {
    
	username := c.Params("username")
	password := c.Params("password")
	
    var login []entities.Login

    result := config.Database.Find(&login, "username = ? and password = ?", username, password)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&login)
}