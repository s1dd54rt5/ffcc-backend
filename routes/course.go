package routes

import (
	"context"
	"net/http"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCourses(c *gin.Context) {
	var ctx context.Context
	var courseList models.CourseList
	collection := db.GetDbCollection("courses-list")
	err := collection.FindOne(ctx, bson.D{{}}).Decode(&courseList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "No courses available"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Data sent succesfully",
		"courses": courseList.Courses,
	})
}
