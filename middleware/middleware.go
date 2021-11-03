package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DummyAuthorizer(c *fiber.Ctx) error {
	fmt.Println("[DummyAuthorizer] --> Authorized!")
	c.Next()
	return nil
}
