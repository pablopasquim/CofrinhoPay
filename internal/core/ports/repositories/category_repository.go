package repositories

import "github.com/pablopasquim/CofrinhoPay/internal/core/domain"

type CategoryRepository interface {
	Create(category *domain.Category) error
	GetByID(id uint) (*domain.Category, error)
	GetAll() ([]*domain.Category, error)
	GetByType(categoryType string) ([]*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id uint) error
}
