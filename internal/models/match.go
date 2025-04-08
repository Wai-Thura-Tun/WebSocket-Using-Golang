package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Match struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User1ID primitive.ObjectID `json:"user1_id" bson:"user1_id"`
	User2ID primitive.ObjectID `json:"user2_id" bson:"user2_id"`
}
