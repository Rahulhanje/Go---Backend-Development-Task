package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/student/go-user-api/internal/logger"
	"go.uber.org/zap"
)

// RequestID middleware adds X-Request-ID to each request
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("X-Request-ID", requestID)
		c.Locals("requestId", requestID)
		return c.Next()
	}
}

// RequestLogger middleware logs request details and duration
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Get request ID from locals
		requestID, _ := c.Locals("requestId").(string)

		// Log request
		logger.Log.Info("Request started",
			zap.String("requestId", requestID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
		)

		// Process request
		err := c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log response
		logger.Log.Info("Request completed",
			zap.String("requestId", requestID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)

		return err
	}
}
