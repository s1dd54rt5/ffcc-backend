package models

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type Course struct {
	UUID    uuid.UUID `json:"uuid" bson:"uuid"`
	Code    string    `json:"code" bson:"code"`
	Credits int       `json:"credits" bson:"credits"`
	Faculty string    `json:"faculty" bson:"faculty"`
	Owner   string    `json:"owner" bson:"owner"`
	Room    string    `json:"room" bson:"room"`
	Slot    string    `json:"slot" bson:"slot"`
	Title   string    `json:"title" bson:"title"`
	Type    string    `json:"type" bson:"type"`
}
