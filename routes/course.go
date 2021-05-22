package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/44t4nk1/ffcc-backend/models"
	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	var courses models.CourseList
	jsonFile, err := os.Open("./csv/courses-list.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &courses)
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Data sent succesfully",
		"courses": courses.Courses,
	})
}
