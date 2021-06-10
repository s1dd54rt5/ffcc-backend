package main

import (
	"log"

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
	log.Fatal(router.Run(":8080"))
}
