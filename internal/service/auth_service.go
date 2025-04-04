package service

import (
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/model"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/repository"
)

func GetUserByEmail(email string) (*model.User, error) {
	return repository.GetUserByEmail(email)
}
