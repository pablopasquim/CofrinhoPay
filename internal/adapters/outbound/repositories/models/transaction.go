package models

import (
	"time"

	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"gorm.io/gorm"
)

type TransactionDB struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	AccountID  uint       `json:"account_id"`
	CategoryID uint       `json:"category_id"`
	Amount     float64    `json:"amount"`
	Type       string     `json:"type"`
	Date       time.Time  `json:"date"`
	Note       string     `json:"note"`
	Account    AccountDB  `json:"account" gorm:"foreignKey:AccountID"`
	Category   CategoryDB `json:"category" gorm:"foreignKey:CategoryID"`
	User       UserDB     `json:"user" gorm:"foreignKey:UserID"`
}

func (t *TransactionDB) ToDomain() *domain.Transaction {
	return &domain.Transaction{
		ID:         t.ID,
		UserID:     t.UserID,
		AccountID:  t.AccountID,
		CategoryID: t.CategoryID,
		Amount:     t.Amount,
		Type:       t.Type,
		Date:       t.Date,
		Note:       t.Note,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}

func TransactionFromDomain(d *domain.Transaction) *TransactionDB {
	return &TransactionDB{
		UserID:     d.UserID,
		AccountID:  d.AccountID,
		CategoryID: d.CategoryID,
		Amount:     d.Amount,
		Type:       d.Type,
		Date:       d.Date,
		Note:       d.Note,
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
	}
}
