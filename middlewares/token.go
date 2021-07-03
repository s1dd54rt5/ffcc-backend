package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(uuid string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uuid"] = uuid
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func IsAuth(endpoint func(c *gin.Context, token *jwt.Token)) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header["Authorization"][0]
		tokenString := reqToken[7:]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error in JWT", "error": true})
			fmt.Println(err)
		} else {
			if token.Valid && token.Claims.(jwt.MapClaims)["authorized"].(bool) {
				endpoint(c, token)
			}
		}
	}
}
