package routes

import (
	"gin-api/internal/handlers"
	"gin-api/internal/middlewares"
	"gin-api/internal/repositories"
	"gin-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	RoleGuide = "guide"
	RoleAdmin = "admin"
)

func RegisterUserRoutes(r *gin.Engine, dbpool *pgxpool.Pool, authRepo *repositories.AuthRepository) {
	userRepo := repositories.NewUserRepository(dbpool)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/register", userHandler.CreateUser)
	r.POST("/login", userHandler.LogIn)
	r.POST("/forgot-password", userHandler.ForgotPassword)
	r.POST("/reset-password/:resetToken", userHandler.ResetPassword)

	protected := r.Group("/")
	protected.Use(middlewares.Protect(authRepo))

	protected.POST("/update-password", userHandler.UpdatePassword)

	userRoutes := protected.Group("/users")
	userRoutes.Use(middlewares.RestrictTo(RoleGuide, RoleAdmin))
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
