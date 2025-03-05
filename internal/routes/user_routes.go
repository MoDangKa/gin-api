package routes

import (
	"gin-api/internal/controllers"
	"gin-api/internal/middlewares"
	"gin-api/internal/repositories"
	"gin-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRoutes(r *gin.Engine, dbpool *pgxpool.Pool) {
	userRepo := repositories.NewUserRepository(dbpool)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r.POST("/register", userController.CreateUser)

	userRoutes := r.Group("/users")
	userRoutes.Use(middlewares.Protect(), middlewares.RestrictTo("guide", "admin"))
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
