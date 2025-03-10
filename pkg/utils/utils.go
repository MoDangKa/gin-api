package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gin-api/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func GetLogFilename() string {
	currentTime := time.Now().Format("2006-01-02")
	return fmt.Sprintf("./logs/%s.log", currentTime)
}

func CreatePasswordResetToken(user *models.User) (string, error) {
	resetToken, err := GenerateToken(32)
	if err != nil {
		return "", fmt.Errorf("failed to generate password reset token: %w", err)
	}

	hashedToken, err := HashToken(resetToken)
	if err != nil {
		return "", fmt.Errorf("failed to hash password reset token: %w", err)
	}

	user.PasswordResetToken = hashedToken
	user.PasswordResetExpires = time.Now().Add(10 * time.Minute)

	log.Printf("Password reset token successfully created for user ID: %d", user.ID)

	return resetToken, nil
}

func GenerateToken(tokenLength int) (string, error) {
	tokenBytes := make([]byte, tokenLength)

	if _, err := rand.Read(tokenBytes); err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	token := hex.EncodeToString(tokenBytes)

	return token, nil
}

func GetResetURL(req *http.Request, resetToken string) string {
	protocol := "http"
	if req.TLS != nil {
		protocol = "https"
	}

	host := req.Host

	resetURL := fmt.Sprintf("%s://%s/reset-password/%s", protocol, host, resetToken)
	return resetURL
}

func GetClaims(c *gin.Context) (jwt.MapClaims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return nil, fmt.Errorf("authorization header missing")
	}

	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return mapClaims, nil
}
