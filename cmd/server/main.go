package main

import (
	"context"
	"log"
	"os"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	log.Print(os.Getenv("APP_ENV"))

	// Connect to MongoDB and Redis
	config.MongoClient = config.ConnectMongoDB()
	config.RedisClient = config.ConnectRedis()

	// Test MongoDB and Redis connection
	log.Print(os.Getenv("MONGO_URI"))
	log.Print(os.Getenv("REDIS_URL"))
	if err := config.MongoClient.Ping(context.Background(), nil); err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	if err := config.RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	app := fiber.New()

	handler.SetupRoutes(app)

	app.Listen(":8080")
}
