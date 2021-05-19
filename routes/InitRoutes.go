package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "Lol")
	})
}
