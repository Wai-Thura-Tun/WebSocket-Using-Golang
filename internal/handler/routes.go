package handler

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	userHandler := NewUserHandler()

	app.Get("/", userHandler.CreateUser)
}
