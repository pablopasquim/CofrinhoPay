package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID     uint     `json:"user_id"`
	AccountID  uint     `json:"account_id"`
	CategoryID uint     `json:"category_id"`
	Amount     float64  `json:"amount"`
	Type       string   `json:"type"`
	Date       string   `json:"date"`
	Note       string   `json:"note"`
	Account    Account  `json:"account" gorm:"foreignKey:AccountID"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	User       User     `json:"user" gorm:"foreignKey:UserID"`
}
