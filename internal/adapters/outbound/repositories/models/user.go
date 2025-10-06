package models

import (
	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	Username string      `json:"username"`
	Email    string      `json:"email" gorm:"uniqueIndex"`
	Password string      `json:"-" gorm:"not null"`
	Accounts []AccountDB `json:"accounts" gorm:"foreignKey:UserID"`
}

func (u *UserDB) ToDomain() *domain.User {
	return &domain.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserFromDomain(d *domain.User) *UserDB {
	return &UserDB{
		Username: d.Username,
		Email:    d.Email,
		Password: d.Password,
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
	}
}
