package config

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient    *mongo.Client
	RedisClient    *redis.Client
	UserCollection *mongo.Collection
	once           sync.Once
	ctx            context.Context
	cancel         context.CancelFunc
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectMongoDB() {
	once.Do(func() {
		ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			log.Fatal("❌ MONGO_URI is not set in environment variables")
		}

		clientOptions := options.Client().ApplyURI(uri).SetConnectTimeout(10 * time.Second)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("❌ Failed to connect to MongoDB:", err)
		}

		// Verify connection
		if err := client.Ping(ctx, nil); err != nil {
			log.Fatal("❌ MongoDB connection failed:", err)
		}

		log.Println("✅ Connected to MongoDB")

		MongoClient = client
		UserCollection = client.Database("test_db").Collection("users")
	})
}

func ConnectRedis() {
	once.Do(func() {
		addr := os.Getenv("REDIS_URL")
		if addr == "" {
			log.Fatal("❌ REDIS_URL is not set in environment variables")
		}

		RedisClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
		})

		// Ping Redis to verify connection
		_, err := RedisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal("❌ Failed to connect to Redis:", err)
		}

		log.Println("✅ Connected to Redis")
	})
}

func GracefulShutdown() {
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal
	<-signChan
	log.Println("🛑 Shutting down gracefully...")

	// Use context for graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Close MongoDB
	if err := MongoClient.Disconnect(shutdownCtx); err != nil {
		log.Println("❌ Error closing MongoDB:", err)
	} else {
		log.Println("✅ MongoDB connection closed")
	}

	// Close Redis
	if err := RedisClient.Close(); err != nil {
		log.Println("❌ Error closing Redis:", err)
	} else {
		log.Println("✅ Redis connection closed")
	}
	cancel() // cancel any ongoing operations
	os.Exit(0)
}
