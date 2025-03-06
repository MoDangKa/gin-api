package routes

import (
	"gin-api/internal/handlers"
	"gin-api/internal/middlewares"
	"gin-api/internal/repositories"
	"gin-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRoutes(r *gin.Engine, dbpool *pgxpool.Pool, authRepo *repositories.AuthRepository) {
	userRepo := repositories.NewUserRepository(dbpool)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/register", userHandler.CreateUser)
	r.POST("/login", userHandler.LogIn)
	r.POST("/forgot-password", userHandler.ForgotPassword)

	userRoutes := r.Group("/users")
	userRoutes.Use(middlewares.Protect(authRepo), middlewares.RestrictTo("guide", "admin"))
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
