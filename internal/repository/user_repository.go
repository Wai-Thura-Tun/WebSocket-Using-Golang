package repository

import (
	"context"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
	userCollection = config.MongoClient.Database("test_db").Collection("users")
}

func CreateUser(user model.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}
