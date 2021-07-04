package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func RateFaculty(c *gin.Context, token *jwt.Token) {
	var fac models.FacultyRating
	if err := c.ShouldBindJSON(&fac); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid JSON Provided"})
		return
	}
	facFound := false
	userRated := false
	payload := token.Claims.(jwt.MapClaims)["id"].(string)
	var facultyList models.FacultyList
	collection := db.GetDbCollection("faculty-list")
	err := collection.FindOne(ctx, bson.D{{}}).Decode(&facultyList)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "true"})
		log.Fatal(err)
	}
	var faculty models.Faculty
	for _, elem := range facultyList.FacultyList {
		if elem.ID == fac.ID {
			facFound = true
			faculty = elem
			break
		}
	}
	if facFound {
		objID, err := primitive.ObjectIDFromHex(payload)
		if err != nil {
			panic(err)
		}
		for _, elem := range faculty.RatedBy {
			fmt.Println(elem)
			if elem == objID {
				userRated = true
				break
			}
		}
		if userRated {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Already rated this Faculty"})
		} else {
			for elem := range facultyList.FacultyList {
				if facultyList.FacultyList[elem].ID == fac.ID {
					objID, err := primitive.ObjectIDFromHex(payload)
					if err != nil {
						panic(err)
					}
					facultyList.FacultyList[elem].RatedBy = append(facultyList.FacultyList[elem].RatedBy, objID)
					facultyList.FacultyList[elem].Reviews = facultyList.FacultyList[elem].Reviews + 1
					facultyList.FacultyList[elem].Rating = (facultyList.FacultyList[elem].Rating + fac.Rating) / float64(facultyList.FacultyList[elem].Reviews)
					break
				}
			}
			_, err := collection.DeleteMany(ctx, bson.D{{}})
			if err != nil {
				log.Fatal(err)
			}
			_, err = collection.InsertOne(ctx, facultyList)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{"error": false, "message": "Faculty rated successfully"})
			c.JSON(http.StatusOK, facultyList)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Incorrect Faculty ID"})
	}
}
