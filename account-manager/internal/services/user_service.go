package services

import (
	"assessment/account-manager/internal/models"
	"assessment/account-manager/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) Register(user *models.User) error {
	return us.userRepository.Create(user)
}

func (us *UserService) Login(username string) (*models.User, error) {
	return us.userRepository.FindUser(username)
}
