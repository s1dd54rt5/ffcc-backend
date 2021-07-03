package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/middlewares"
	"github.com/44t4nk1/ffcc-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func Signup(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid JSON Provided"})
		return
	}
	filter := bson.D{primitive.E{Key: "email", Value: u.Email}}
	var result models.User
	collection := db.GetDbCollection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		user := models.User{
			ID:       primitive.NewObjectID(),
			Email:    u.Email,
			Password: hashAndSalt(([]byte(u.Password))),
		}
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"error":   false,
			"message": "Account created succesfully",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
	}
}

func Login(c *gin.Context) {
	var result, u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON Provided")
		return
	}
	filter := bson.D{primitive.E{Key: "email", Value: u.Email}}
	collection := db.GetDbCollection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User does not exist"})
		return
	}
	if comparePasswords(result.Password, []byte(u.Password)) {
		tokenString := result.ID.String()
		token, err := middlewares.CreateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error in creating token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"error":   false,
			"message": "Login succesfull",
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Incorrect password",
		})
	}
}
