package main

import (
	"log"
	"os"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/handlers"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/ws"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	log.Print(os.Getenv("APP_ENV"))

	config.ConnectMongoDB()
	config.ConnectRedis()

	//  Initilize GateWay
	gateway := ws.NewGateWay()
	go gateway.Run()

	app := fiber.New()

	handlers.SetupRoutes(app, gateway)

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
