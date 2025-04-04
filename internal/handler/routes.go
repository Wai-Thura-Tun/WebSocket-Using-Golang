package handler

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	userHandler := NewUserHandler()
	authHandler := NewAuthHandler()

	app.Post("/register", userHandler.CreateUser)
	app.Post("/login", authHandler.Login)
}
