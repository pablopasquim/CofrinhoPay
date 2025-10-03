package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:CategoryID"`
}
