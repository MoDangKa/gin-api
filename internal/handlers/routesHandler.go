package routesHandler

import (
	usersHandler "gin-api/internal/handlers/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *usersHandler.UserHandler) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
}
