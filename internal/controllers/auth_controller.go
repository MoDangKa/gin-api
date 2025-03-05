package controllers

import (
	"gin-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (h *AuthController) LogIn(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	result, err := h.authService.LogIn(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"user":  result.User,
		"token": result.Token,
	}

	c.JSON(http.StatusOK, response)
}

// func GetUserByEmailAndActive(db *sqlx.DB, email string) (*User, error) {
// 	var user User
// 	query := `SELECT * FROM public.users WHERE email = $1 AND active = TRUE`
// 	err := db.Get(&user, query, email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
