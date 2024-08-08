package models

import (
	"gorm.io/gorm"
)

type PaymentHistory struct {
	gorm.Model
	UserID        string  `json:"user_id"`
	TransactionID uint    `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"` // "send" or "withdraw"
}
