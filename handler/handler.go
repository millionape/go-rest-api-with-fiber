package handler

import (
	"github.com/gofiber/fiber/v2"
)

// GetHelloworld
func GetHelloworld(c *fiber.Ctx) error {

	// Return in JSON format
	c.JSON(&fiber.Map{
		"success": true,
		"text":    "Hello world",
	})
	return nil
}
