package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashToken(token string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(token)); err != nil {
		return "", fmt.Errorf("error hashing token: %w", err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
