package routes

import (
	"gin-api/internal/controllers"
	"gin-api/internal/middlewares"
	"gin-api/internal/repositories"
	"gin-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRoutes(router *gin.Engine, dbpool *pgxpool.Pool) {
	userRepo := repositories.NewUserRepository(dbpool)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userRoutes := router.Group("/users")
	userRoutes.POST("/", userController.CreateUser)
	userRoutes.POST("/login", userController.LogIn)

	protectedRoutes := userRoutes.Group("/")
	protectedRoutes.Use(middlewares.Auth(), middlewares.IsAdmin())
	{
		protectedRoutes.GET("/", userController.GetAllUsers)
		protectedRoutes.GET("/:id", userController.GetUserByID)
		protectedRoutes.PUT("/:id", userController.UpdateUser)
		protectedRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
