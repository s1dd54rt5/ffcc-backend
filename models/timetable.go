package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TimeTable struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Slots   [60]bool           `json:"slots" bson:"slots"`
	Courses []TimeTableCourse  `json:"courses" bson:"courses"`
}

type TimeTableCourse struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Code    string             `json:"code" bson:"code"`
	Title   string             `json:"title" bson:"title"`
	Credits int                `json:"credits" bson:"credits"`
	Type    string             `json:"type" bson:"type"`
	Slot    []int              `json:"slot" bson:"slot"`
}
