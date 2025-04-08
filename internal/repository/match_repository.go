package repository

import (
	"context"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMatch(user1ID string, user2ID string) error {
	var match models.Match
	user1, err := primitive.ObjectIDFromHex(user1ID)
	if err != nil {
		return err
	}

	user2, err := primitive.ObjectIDFromHex(user2ID)
	if err != nil {
		return err
	}

	match.User1ID = user1
	match.User2ID = user2
	_, err = config.MatchCollection.InsertOne(context.Background(), match)
	return err
}

func GetMatches() ([]models.Match, error) {
	cursor, err := config.MatchCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var matches []models.Match
	if err = cursor.All(context.Background(), &matches); err != nil {
		return nil, err
	}
	return matches, nil
}
