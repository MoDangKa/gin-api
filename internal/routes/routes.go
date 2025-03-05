package routes

import (
	"gin-api/internal/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(r *gin.Engine, dbpool *pgxpool.Pool) {
	r.Use(middlewares.RateLimiter(100, time.Hour))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	RegisterAuthRoutes(r, dbpool)
	RegisterUserRoutes(r, dbpool)
}
