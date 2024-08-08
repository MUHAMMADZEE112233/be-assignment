package controllers

import (
	"assessment/payment-manager/internal/models"
	"assessment/payment-manager/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{TransactionService: transactionService}
}

func (tc *TransactionController) Send(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TransactionService.ProcessTransaction(transaction, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction processed successfully"})
}

func (tc *TransactionController) Withdraw(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TransactionService.ProcessTransaction(transaction, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction processed successfully"})
}

func (tc *TransactionController) GetTransactions(c *gin.Context) {
	transactions, err := tc.TransactionService.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
