package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the request has a valid authorization token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		// TODO: Validate the token (e.g., using JWT)
		isValid := validateToken(token)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// If token is valid, proceed to the next handler
		c.Next()
	}
}

// validateToken validates the token (this is a placeholder function)
func validateToken(token string) bool {
	// Implement token validation logic here (e.g., using JWT)
	return token == "valid-token" // Replace with actual validation logic
}
