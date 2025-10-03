package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Password string    `json:"password"`
	Accounts []Account `json:"accounts" gorm:"foreignKey:UserID"`
}
