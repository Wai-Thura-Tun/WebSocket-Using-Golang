package service

import (
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/model"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/repository"
)

func CreateUser(user model.User) error {
	return repository.CreateUser(user)
}
