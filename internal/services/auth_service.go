package services

import (
	"gin-api/internal/models"
	"gin-api/internal/repositories"
)

type AuthService struct {
	authRepo *repositories.AuthRepository
}

func NewAuthService(authRepo *repositories.AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) LogIn(email string, password string) (*models.Auth, error) {
	return s.authRepo.LogIn(email, password)
}
