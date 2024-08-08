package repository

import (
	"assessment/payment-manager/internal/models"

	"gorm.io/gorm"
)

type PaymentHistoryRepository struct {
	Database *gorm.DB
}

func NewPaymentHistoryRepository(database *gorm.DB) *PaymentHistoryRepository {
	return &PaymentHistoryRepository{Database: database}
}

func (phr *PaymentHistoryRepository) CreatePaymentHistory(history models.PaymentHistory) error {
	return phr.Database.Create(&history).Error
}
