package main

import (
	"assessment/account-manager/internal/controllers"
	"assessment/account-manager/internal/middleware"
	"assessment/account-manager/internal/repository"
	"assessment/account-manager/internal/services"
	"assessment/account-manager/pkg/db"
	"assessment/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	cfg := config.LoadConfig()
	db, err := db.Init(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Set up the router
	router := gin.Default()

	// Initialize repositories
	accountRepo := repository.NewAccountRepository(db.DB)

	userRepo := repository.NewUserRepository(db.DB)
	// Initialize services
	accountService := services.NewAccountService(accountRepo)

	userService := services.NewUserService(userRepo)
	// Initialize controllers
	accountController := controllers.NewAccountController(accountService)
	userController := controllers.NewUserController(userService)
	// User routes
	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)

	// Account routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/accounts", accountController.CreateAccount)
		authorized.GET("/accounts", accountController.GetAccounts)
		authorized.GET("/accounts/:accountId/transactions", accountController.GetTransactions)
	}
	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
