package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(maxRequests int, duration time.Duration) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(duration), maxRequests)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests from this IP, please try again in an hour!",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
