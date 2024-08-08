package main

import (
	"assessment/config"
	"assessment/payment-manager/internal/controllers"
	"assessment/payment-manager/internal/middleware"
	"assessment/payment-manager/internal/repository"
	"assessment/payment-manager/internal/services"
	"assessment/payment-manager/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.LoadConfig()
	database := db.InitDB(cfg)
	transactionRepo := repository.NewTransactionRepository(database.DB)
	paymentHistoryRepo := repository.NewPaymentHistoryRepository(database.DB)
	accountRepo := repository.NewAccountRepository(database.DB)
	transactionService := services.NewTransactionService(transactionRepo, paymentHistoryRepo, accountRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/send", transactionController.Send)
		authorized.POST("/withdraw", transactionController.Withdraw)
		authorized.GET("/transactions", transactionController.GetTransactions)
	}
	r.Run(":8081")
}
