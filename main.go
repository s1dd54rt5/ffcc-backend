package main

import (
	"log"

	"github.com/44t4nk1/ffcc-backend/routes"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	routes.InitRoutes(router)
	log.Fatal(router.Run(":8080"))
}
