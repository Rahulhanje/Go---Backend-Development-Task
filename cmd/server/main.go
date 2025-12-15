package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/student/go-user-api/config"
	"github.com/student/go-user-api/db/sqlc"
	"github.com/student/go-user-api/internal/handler"
	"github.com/student/go-user-api/internal/logger"
	"github.com/student/go-user-api/internal/repository"
	"github.com/student/go-user-api/internal/routes"
	"github.com/student/go-user-api/internal/service"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Log.Info("Starting application...")

	// Load configuration
	cfg := config.Load()
	logger.Log.Info("Configuration loaded")

	// Connect to database
	db, err := config.ConnectDB(cfg)
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()
	logger.Log.Info("Connected to database")

	// Initialize layers
	queries := sqlc.New(db)
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Log.Error("Unhandled error", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		},
	})

	// Add recover middleware to prevent panics
	app.Use(recover.New())

	// Setup routes
	routes.SetupRoutes(app, userHandler)

	// Start server
	serverAddr := ":" + cfg.ServerPort
	logger.Log.Info("Server starting", zap.String("address", serverAddr))

	if err := app.Listen(serverAddr); err != nil {
		logger.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
