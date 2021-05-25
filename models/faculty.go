package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Faculty struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Faculty string             `json:"faculty" bson:"faculty"`
	Room    string             `json:"room" bson:"room"`
	Slot    string             `json:"slot" bson:"slot"`
	Rating  float64            `json:"rating" bson:"rating"`
	Reviews int                `json:"reviews" bson:"reviews"`
}
