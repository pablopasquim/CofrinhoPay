package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	Balance      float64       `json:"balance"`
	UserID       uint          `json:"user_id"`
	User         User          `json:"user" gorm:"foreignKey:UserID"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:AccountID"`
}
