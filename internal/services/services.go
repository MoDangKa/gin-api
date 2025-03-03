package services

import (
	"gin-api/internal/models"
	"gin-api/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}
