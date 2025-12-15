package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/student/go-user-api/internal/handler"
	"github.com/student/go-user-api/internal/middleware"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	// Apply global middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	// User routes
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Get("/users", userHandler.ListUsers)
}
