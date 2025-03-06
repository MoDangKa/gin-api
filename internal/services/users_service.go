package services

import (
	"gin-api/internal/models"
	"gin-api/internal/repositories"
	"net/http"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}

func (s *UserService) LogIn(email string, password string) (*models.UserWithToken, error) {
	return s.userRepo.LogIn(email, password)
}

func (s *UserService) ForgotPassword(req *http.Request, email string) error {
	return s.userRepo.ForgotPassword(req, email)
}
