package routes

import (
	"gin-api/internal/handlers"
	"gin-api/internal/middlewares"
	"gin-api/internal/repositories"
	"gin-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRoutes(router *gin.Engine, dbpool *pgxpool.Pool) {
	userRepo := repositories.NewUserRepository(dbpool)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Define user routes
	userRoutes := router.Group("/users")

	// Public route: Get all users
	userRoutes.GET("/", userHandler.GetAllUsers)

	// Protected routes: Apply AuthMiddleware
	protectedRoutes := userRoutes.Group("/")
	protectedRoutes.Use(middlewares.AuthMiddleware())
	{
		protectedRoutes.POST("/", userHandler.CreateUser)
		protectedRoutes.GET("/:id", userHandler.GetUserByID)
		protectedRoutes.PUT("/:id", userHandler.UpdateUser)
		protectedRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
