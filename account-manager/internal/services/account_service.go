package services

import (
	"assessment/account-manager/internal/models"
	"assessment/account-manager/internal/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (as *AccountService) CreateAccount(account *models.Account) error {
	return as.accountRepository.Create(account)
}

func (as *AccountService) GetAccounts(userID uint) ([]models.Account, error) {
	return as.accountRepository.FindByUserID(userID)
}

func (as *AccountService) GetTransactions(accountID string) ([]models.PaymentHistory, error) {
	return as.accountRepository.FindTransactionsByAccountID(accountID)
}
