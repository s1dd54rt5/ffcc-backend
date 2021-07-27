package main

import (
	"log"
	"os"
	"time"

	"github.com/44t4nk1/ffcc-backend/db"
	"github.com/44t4nk1/ffcc-backend/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

var (
	router = gin.Default()
)

func init() {
	db.InitialiseDb()
	rate, err := limiter.NewRateFromFormatted("1-M")
	if err != nil {
		log.Fatal(err)
		return
	}
	store := memory.NewStore()
	rateLimiter := mgin.NewMiddleware(limiter.New(store, rate))
	router.Use(
		cors.Middleware(
			cors.Config{
				Origins:         "*",
				Methods:         "GET, PUT, POST, DELETE",
				RequestHeaders:  "Origin, Authorization, Content-Type",
				ExposedHeaders:  "",
				MaxAge:          50 * time.Second,
				Credentials:     true,
				ValidateHeaders: false,
			}))
	router.Use(rateLimiter)
	routes.InitRoutes(router)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}
