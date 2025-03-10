package middlewares

import (
	"fmt"
	"gin-api/internal/repositories"
	"gin-api/pkg/utils"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Protect(authRepo *repositories.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		email := claims["email"].(string)
		isActive, err := authRepo.IsActive(email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error checking user active status"})
			c.Abort()
			return
		}

		if !isActive {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not active"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func RestrictTo(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := utils.GetClaims(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: role privileges required"})
			c.Abort()
			return
		}

		if !slices.Contains(roles, role) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: you do not have permission to perform this action"})
			c.Abort()
			return
		}

		c.Next()
	}
}
