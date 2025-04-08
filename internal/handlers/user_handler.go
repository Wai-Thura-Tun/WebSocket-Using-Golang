package handlers

import (
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/models"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	err := services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
