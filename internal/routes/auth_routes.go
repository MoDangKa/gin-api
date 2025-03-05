package routes

import (
	"gin-api/internal/controllers"
	"gin-api/internal/repositories"
	"gin-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterAuthRoutes(r *gin.Engine, dbpool *pgxpool.Pool) {
	authRepo := repositories.NewAuthRepository(dbpool)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	r.POST("/login", authController.LogIn)
}
