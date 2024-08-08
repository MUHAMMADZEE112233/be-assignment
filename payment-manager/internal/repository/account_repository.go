package repository

import (
	"assessment/payment-manager/internal/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (ar *AccountRepository) GetAccountByUserID(userID uint) (*models.Account, error) {
	var account models.Account
	if err := ar.DB.Where("id = ?", userID).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (ar *AccountRepository) UpdateAccount(account *models.Account) error {
	return ar.DB.Save(account).Error
}
