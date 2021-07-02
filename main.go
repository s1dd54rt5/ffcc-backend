package main

import (
	"fmt"
	"log"
	"os"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/routes"
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
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
