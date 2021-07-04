package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Faculty struct {
	ID      primitive.ObjectID   `json:"_id" bson:"_id"`
	Name    string               `json:"faculty" bson:"faculty"`
	Rating  float64              `json:"rating" bson:"rating"`
	Reviews int                  `json:"reviews" bson:"reviews"`
	RatedBy []primitive.ObjectID `json:"ratedby" bson:"ratedby"`
}

type FacultyList struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	FacultyList []Faculty          `json:"faculty" bson:"faculty"`
}

type FacultyRating struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Rating float64            `json:"rating" bson:"rating"`
}
