package repository

import (
	"context"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := config.UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
