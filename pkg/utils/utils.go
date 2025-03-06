package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gin-api/internal/models"
	"net/http"
	"time"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func GetLogFilename() string {
	currentTime := time.Now().Format("2006-01-02")
	return fmt.Sprintf("./logs/%s.log", currentTime)
}

func CreatePasswordResetToken(user *models.User) (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}
	resetToken := hex.EncodeToString(tokenBytes)

	hash := sha256.New()
	_, err = hash.Write([]byte(resetToken))
	if err != nil {
		return "", fmt.Errorf("error hashing token: %w", err)
	}
	hashedToken := hex.EncodeToString(hash.Sum(nil))

	user.PasswordResetToken = hashedToken
	user.PasswordResetExpires = time.Now().Add(10 * time.Minute)

	fmt.Printf("Reset Token: %s\nHashed Token: %s\n", resetToken, user.PasswordResetToken)

	return resetToken, nil
}

func DecodeResetToken(user *models.User, resetToken string) (bool, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(resetToken))
	if err != nil {
		return false, fmt.Errorf("error hashing token: %w", err)
	}
	hashedToken := hex.EncodeToString(hash.Sum(nil))

	if hashedToken != user.PasswordResetToken {
		return false, fmt.Errorf("invalid or expired token")
	}

	if time.Now().After(user.PasswordResetExpires) {
		return false, fmt.Errorf("token has expired")
	}

	return true, nil
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
