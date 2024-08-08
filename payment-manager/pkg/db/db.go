package db

import (
	"assessment/config"
	"assessment/payment-manager/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func InitDB(cfg *config.Config) *Database {
	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=" + cfg.DBSSLMode

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&models.Transaction{}, &models.PaymentHistory{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("Database connected and migrated")
	return &Database{DB: db}
}
