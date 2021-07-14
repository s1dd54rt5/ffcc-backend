package main

import (
	"log"
	"os"
	"time"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var (
	router = gin.Default()
)

func init() {
	db.InitialiseDb()
	router.Use(cors.Middleware(
		cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			Credentials:     true,
			ValidateHeaders: false,
		}))
	routes.InitRoutes(router)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}
