package routes

import (
	"log"
	"net/http"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFaculty(c *gin.Context) {
	var facultyList models.FacultyList
	collection := db.GetDbCollection("faculty-list")
	err := collection.FindOne(ctx, bson.D{{}}).Decode(&facultyList)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"error":     false,
		"message":   "Data sent succesfully",
		"faculties": facultyList.FacultyList,
	})
}
