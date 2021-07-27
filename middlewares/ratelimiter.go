package middlewares

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimitHandler(c *gin.Context) {
	c.JSON(http.StatusTooManyRequests, gin.H{"error": true, "message": "Rate Limit reached"})
}

func CustomMiddleware() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("5-S")
	if err != nil {
		fmt.Println("Error!")
	}

	store := memory.NewStore()

	limit := &limiter.Limiter{
		Rate:  rate,
		Store: store,
	}
	middleware := &mgin.Middleware{
		Limiter:        limit,
		OnError:        mgin.DefaultErrorHandler,
		OnLimitReached: RateLimitHandler,
		KeyGetter:      mgin.DefaultKeyGetter,
		ExcludedKey:    nil,
	}

	return func(ctx *gin.Context) {
		middleware.Handle(ctx)
	}
}
