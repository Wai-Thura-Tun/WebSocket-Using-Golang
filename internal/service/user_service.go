package service

import (
	"errors"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/model"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/repository"
)

func CreateUser(user model.User) error {
	existingUser, _ := repository.GetUserByID(user.ID.Hex())
	if existingUser != nil {
		return errors.New("user already exists")
	}
	return repository.CreateUser(user)
}

func GetUserByID(userID string) (*model.User, error) {
	return repository.GetUserByID(userID)
}
