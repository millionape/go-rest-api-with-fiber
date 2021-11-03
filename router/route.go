package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/millionape/go-rest-api-with-fiber/handler"
	"github.com/millionape/go-rest-api-with-fiber/handler/user"
	"github.com/millionape/go-rest-api-with-fiber/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api", middleware.DummyAuthorizer)
	// routes
	api.Get("/", handler.GetHelloworld)
	api.Get("/user", user.GetUser)

}
