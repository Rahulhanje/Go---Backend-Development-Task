package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/student/go-user-api/internal/logger"
	"github.com/student/go-user-api/internal/models"
	"github.com/student/go-user-api/internal/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	if err := models.ValidateStruct(req); err != nil {
		logger.Log.Error("Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
	}

	user, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to create user",
			Message: err.Error(),
		})
	}

	logger.Log.Info("User created", zap.Int32("id", user.ID))
	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid user ID",
			Message: err.Error(),
		})
	}

	user, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "User not found",
			})
		}
		logger.Log.Error("Failed to get user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to get user",
			Message: err.Error(),
		})
	}

	return c.JSON(user)
}

// UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid user ID",
			Message: err.Error(),
		})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	if err := models.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
	}

	user, err := h.service.UpdateUser(c.Context(), int32(id), req)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "User not found",
			})
		}
		logger.Log.Error("Failed to update user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to update user",
			Message: err.Error(),
		})
	}

	logger.Log.Info("User updated", zap.Int32("id", user.ID))
	return c.JSON(user)
}

// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid user ID",
			Message: err.Error(),
		})
	}

	if err := h.service.DeleteUser(c.Context(), int32(id)); err != nil {
		logger.Log.Error("Failed to delete user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to delete user",
			Message: err.Error(),
		})
	}

	logger.Log.Info("User deleted", zap.Int("id", id))
	return c.SendStatus(fiber.StatusNoContent)
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		logger.Log.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to list users",
			Message: err.Error(),
		})
	}

	return c.JSON(users)
}
