package services

import (
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/models"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/repository"
)

func CreateMatch(user1ID string, user2ID string) error {
	return repository.CreateMatch(user1ID, user2ID)
}

func GetMatches() ([]models.Match, error) {
	return repository.GetMatches()
}
