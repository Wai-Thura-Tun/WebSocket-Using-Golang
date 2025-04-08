package services

import (
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/models"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/repository"
)

func GetUserByEmail(email string) (*models.User, error) {
	return repository.GetUserByEmail(email)
}
