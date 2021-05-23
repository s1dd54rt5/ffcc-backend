package routes

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx context.Context
)

func LoadCsv(c *gin.Context) {
	csvfile, err := os.Open("./csv/ffcc.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvfile)
	var course models.Course
	collection := db.GetDbCollection("courses")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		credit, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		course = models.Course{
			ID:      primitive.NewObjectID(),
			Code:    record[0],
			Credits: credit,
			Faculty: record[2],
			Owner:   record[3],
			Room:    record[4],
			Slot:    record[5],
			Title:   record[6],
			Type:    record[7],
			Rating:  0.0,
			Reviews: 0,
		}
		_, err = collection.InsertOne(ctx, course)
		if err != nil {
			log.Fatal(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added data"})
}

func CourseList(c *gin.Context) {
	var ctx context.Context
	var courses []models.CourseItem
	findOptions := options.Find()
	findOptions.SetLimit(6000)
	collection := db.GetDbCollection("courses")
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var elem models.CourseItem
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		if !contains(&courses, &elem) {
			elem.ID = primitive.NewObjectID()
			courses = append(courses, elem)
		}
	}
	listCollection := db.GetDbCollection("courses-list")
	listCourses := models.CourseList{
		ID:      primitive.NewObjectID(),
		Courses: courses,
	}
	_, err = listCollection.InsertOne(ctx, listCourses)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Data added succesfully",
	})
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(ctx)
}

func contains(courses *[]models.CourseItem, course *models.CourseItem) bool {
	flag := false
	for _, elem := range *courses {
		if elem.Code == course.Code && elem.Type == course.Type {
			flag = true
			break
		}
	}
	return flag
}
