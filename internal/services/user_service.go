package services

import (
	"errors"
	"log"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/models"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/repository"
)

func CreateUser(user models.User) error {
	existingUser, _ := repository.GetUserByID(user.ID.Hex())

	if existingUser != nil {
		log.Print("user already exist")
		return errors.New("user already exists")
	}
	return repository.CreateUser(user)
}

func GetUserByID(userID string) (*models.User, error) {
	return repository.GetUserByID(userID)
}
