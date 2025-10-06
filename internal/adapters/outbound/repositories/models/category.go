package models

import (
	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"gorm.io/gorm"
)

type CategoryDB struct {
	gorm.Model
	Name         string          `json:"name"`
	Type         string          `json:"type"`
	Transactions []TransactionDB `json:"transactions" gorm:"foreignKey:CategoryID"`
}

func (c *CategoryDB) ToDomain() *domain.Category {
	return &domain.Category{
		ID:        c.ID,
		Name:      c.Name,
		Type:      c.Type,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func CategoryFromDomain(d *domain.Category) *CategoryDB {
	return &CategoryDB{
		Name: d.Name,
		Type: d.Type,
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
	}
}
