package handlers

import (
	"log"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/services"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	log.Println("Raw body:", string(c.Body()))
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Payload",
		})
	}

	log.Println("email: ", req.Email)
	log.Println("password: ", req.Password)
	existUser, err := services.GetUserByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	token, err := utils.GenerateToken(existUser.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
