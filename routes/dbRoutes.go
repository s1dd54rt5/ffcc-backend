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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		}
		_, err = collection.InsertOne(ctx, course)
		if err != nil {
			log.Fatal(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added data"})
}
