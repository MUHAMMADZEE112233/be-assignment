package repository

import (
	"assessment/account-manager/internal/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (ar *AccountRepository) Create(account *models.Account) error {
	return ar.DB.Create(account).Error
}

func (ar *AccountRepository) FindByUserID(userID uint) ([]models.Account, error) {
	var accounts []models.Account
	if err := ar.DB.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (ar *AccountRepository) FindTransactionsByAccountID(accountID string) ([]models.PaymentHistory, error) {
	var transactions []models.PaymentHistory
	if err := ar.DB.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
