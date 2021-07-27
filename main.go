package main

import (
	"fmt"
	"log"
	"net/http"
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

func RateLimitHandler(c *gin.Context) {
	c.JSON(http.StatusTooManyRequests, gin.H{"error": true, "message": "Rate Limit reached"})
}

func CustomMiddleware(limiter *limiter.Limiter) gin.HandlerFunc {
	middleware := &mgin.Middleware{
		Limiter:        limiter,
		OnError:        mgin.DefaultErrorHandler,
		OnLimitReached: RateLimitHandler,
		KeyGetter:      mgin.DefaultKeyGetter,
		ExcludedKey:    nil,
	}

	return func(ctx *gin.Context) {
		middleware.Handle(ctx)
	}
}

func init() {
	db.InitialiseDb()
	rate, err := limiter.NewRateFromFormatted("5-S")
	if err != nil {
		fmt.Println("Error!")
	}

	store := memory.NewStore()

	limit := &limiter.Limiter{
		Rate:  rate,
		Store: store,
	}

	rateLimiter := CustomMiddleware(limit)
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
