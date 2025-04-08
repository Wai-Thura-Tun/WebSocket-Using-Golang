package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MediaHandler struct{}

func NewMediaHandler() *MediaHandler {
	return &MediaHandler{}
}

func (h *MediaHandler) UploadMedia(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)
	file, err := c.FormFile("media")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing media file",
		})
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create upload folder",
		})
	}

	fileName := fmt.Sprintf("%s/%d_%s", userID, time.Now().UnixMilli(), file.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	// Save locally
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	url := fmt.Sprintf("http://127.0.0.1:8080/uploads/%s", fileName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"media_url": url,
	})
}
