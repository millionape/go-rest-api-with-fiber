package user

import (
	"github.com/gofiber/fiber/v2"
)

// GetHelloworld
func GetUser(c *fiber.Ctx) error {

	// Return in JSON format
	c.JSON(&fiber.Map{
		"success": true,
		"text":    "User1",
	})
	return nil
}
