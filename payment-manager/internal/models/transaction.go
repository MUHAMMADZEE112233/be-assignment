package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount      float64 `json:"amount"`
	FromAddress uint    `json:"fromAddress"`
	ToAddress   uint    `json:"toAddress"`
	Status      string  `json:"status"`
}
