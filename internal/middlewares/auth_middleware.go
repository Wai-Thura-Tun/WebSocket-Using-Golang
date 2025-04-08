package middlewares

import (
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	userId, err := utils.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	c.Locals("userId", userId)
	return c.Next()
}
