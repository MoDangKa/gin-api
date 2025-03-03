package handlers

import (
	"gin-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *services.UserService
}

func NewHandler(userService *services.UserService) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	r.GET("/health", h.HealthCheck)
	r.GET("/users", h.GetUsers)
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
