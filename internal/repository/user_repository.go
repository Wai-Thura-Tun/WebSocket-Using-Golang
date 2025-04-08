package repository

import (
	"context"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	_, err = config.UserCollection.InsertOne(context.Background(), user)
	return err
}

func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	err = config.UserCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
