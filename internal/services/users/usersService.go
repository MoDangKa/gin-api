package usersService

import (
	usersModel "gin-api/internal/models/users"
	usersRepository "gin-api/internal/repositories/users"
)

type UserService struct {
	userRepo *usersRepository.UserRepository
}

func NewUserService(userRepo *usersRepository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]usersModel.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) CreateUser(user *usersModel.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*usersModel.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *usersModel.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
