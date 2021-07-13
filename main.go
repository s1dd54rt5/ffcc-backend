package main

import (
	"log"
	"os"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func init() {
	db.InitialiseDb()
	routes.InitRoutes(router)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Use(cors.Default())
	log.Fatal(router.Run(":" + port))
}
