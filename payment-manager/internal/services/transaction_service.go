package services

import (
	"assessment/payment-manager/internal/models"
	"assessment/payment-manager/internal/repository"
	"errors"
	"time"
)

type TransactionService struct {
	TransactionRepository    *repository.TransactionRepository
	AccountRepository        *repository.AccountRepository
	PaymentHistoryRepository *repository.PaymentHistoryRepository
}

func NewTransactionService(transactionRepository *repository.TransactionRepository, paymentHistoryRepository *repository.PaymentHistoryRepository, accountRepository *repository.AccountRepository) *TransactionService {
	return &TransactionService{
		TransactionRepository:    transactionRepository,
		AccountRepository:        accountRepository,
		PaymentHistoryRepository: paymentHistoryRepository,
	}
}

func (ts *TransactionService) ProcessTransaction(transaction models.Transaction, isSend bool) error {
	// Start a database transaction
	tx := ts.TransactionRepository.BeginTransaction()

	// Check if sender has sufficient balance
	senderAccount, err := ts.AccountRepository.GetAccountByUserID(transaction.FromAddress)
	if err != nil {
		tx.Rollback()
		return err
	}

	if senderAccount.UserID != transaction.FromAddress {
		tx.Rollback()
		return errors.New("unauthorized transaction: account does not belong to the user")
	}

	if senderAccount.Balance < transaction.Amount {
		tx.Rollback()
		return errors.New("insufficient balance")
	}

	// Deduct amount from sender's account
	senderAccount.Balance -= transaction.Amount
	if err := ts.AccountRepository.UpdateAccount(senderAccount); err != nil {
		tx.Rollback()
		return err
	}

	var receiverAccount *models.Account
	if isSend {
		// Fetch receiver's account
		receiverAccount, err = ts.AccountRepository.GetAccountByUserID(transaction.ToAddress)
		if err != nil {
			tx.Rollback()
			return err
		}

		// Add amount to receiver's account
		receiverAccount.Balance += transaction.Amount
		if err := ts.AccountRepository.UpdateAccount(receiverAccount); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Create the transaction
	transaction.Status = "processing"
	if err := ts.TransactionRepository.CreateTransaction(&transaction); err != nil { // Pass pointer here
		tx.Rollback()
		return err
	}

	// Add payment history for sender
	senderHistory := models.PaymentHistory{
		UserID:        transaction.FromAddress,
		TransactionID: transaction.ID, // Transaction ID should now be set correctly
		Amount:        transaction.Amount,
		Type:          "send",
		AccountID:     senderAccount.ID, // Set AccountID for the sender
	}
	if !isSend {
		senderHistory.Type = "withdraw"
	}
	if err := ts.PaymentHistoryRepository.CreatePaymentHistory(senderHistory); err != nil {
		tx.Rollback()
		return err
	}

	if isSend {
		// Add payment history for receiver
		receiverHistory := models.PaymentHistory{
			UserID:        transaction.ToAddress,
			TransactionID: transaction.ID, // Transaction ID should now be set correctly
			Amount:        transaction.Amount,
			Type:          "receive",
			AccountID:     receiverAccount.ID, // Set AccountID for the receiver
		}
		if err := ts.PaymentHistoryRepository.CreatePaymentHistory(receiverHistory); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	go ts.completeTransaction(transaction)
	return nil
}

func (ts *TransactionService) completeTransaction(transaction models.Transaction) {
	time.Sleep(30 * time.Second)
	transaction.Status = "completed"
	ts.TransactionRepository.UpdateTransactionStatus(transaction)

	// Add balance to receiver's account if it's a send transaction
	account, err := ts.AccountRepository.GetAccountByUserID(transaction.ToAddress)
	if err == nil {
		account.Balance += transaction.Amount
		ts.AccountRepository.UpdateAccount(account)
	}

	// Add payment history
	history := models.PaymentHistory{
		UserID:        transaction.ToAddress,
		TransactionID: transaction.ID,
		Amount:        transaction.Amount,
		Type:          "receive",
	}
	ts.PaymentHistoryRepository.CreatePaymentHistory(history)
}

func (ts *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return ts.TransactionRepository.GetAllTransactions()
}
