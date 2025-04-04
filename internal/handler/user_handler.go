package handler

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	return c.SendString("Create User")
}
