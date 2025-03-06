package routes

import (
	"gin-api/internal/middlewares"
	"gin-api/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(r *gin.Engine, dbpool *pgxpool.Pool) {
	r.Use(middlewares.RateLimiter(100, time.Hour))

	authRepo := repositories.NewAuthRepository(dbpool)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	RegisterUserRoutes(r, dbpool, authRepo)
}
