package repository

import (
	"assessment/account-manager/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) FindUser(username string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
