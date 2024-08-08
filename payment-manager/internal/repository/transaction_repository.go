package repository

import (
	"assessment/payment-manager/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	Database *gorm.DB
}

func NewTransactionRepository(database *gorm.DB) *TransactionRepository {
	return &TransactionRepository{Database: database}
}

func (tr *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	err := tr.Database.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TransactionRepository) UpdateTransactionStatus(transaction models.Transaction) error {
	return tr.Database.Model(&transaction).Update("status", transaction.Status).Error
}

func (tr *TransactionRepository) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := tr.Database.Find(&transactions).Error
	return transactions, err
}

func (tr *TransactionRepository) BeginTransaction() *gorm.DB {
	return tr.Database.Begin()
}
