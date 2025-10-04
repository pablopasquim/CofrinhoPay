package models

import (
	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"gorm.io/gorm"
)

type AccountDB struct {
	gorm.Model
	Name         string          `json:"name"`
	Type         string          `json:"type"`
	Balance      float64         `json:"balance"`
	UserID       uint            `json:"user_id"`
	User         UserDB          `json:"user" gorm:"foreignKey:UserID"`
	Transactions []TransactionDB `json:"transactions" gorm:"foreignKey:AccountID"`
}

func (a *AccountDB) ToDomain() *domain.Account {
	return &domain.Account{
		ID:        a.ID,
		Name:      a.Name,
		Type:      a.Type,
		Balance:   a.Balance,
		UserID:    a.UserID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func AccountFromDomain(d *domain.Account) *AccountDB {
	return &AccountDB{
		Name:    d.Name,
		Type:    d.Type,
		Balance: d.Balance,
		UserID:  d.UserID,
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
	}
}
