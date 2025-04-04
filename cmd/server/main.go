package main

import (
	"log"
	"os"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	log.Print(os.Getenv("APP_ENV"))

	config.ConnectMongoDB()
	config.ConnectRedis()

	app := fiber.New()

	handler.SetupRoutes(app)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	config.GracefulShutdown()

	// Exit the program after shutdown
	log.Println("Graceful shutdown completed")
	os.Exit(0)
}
