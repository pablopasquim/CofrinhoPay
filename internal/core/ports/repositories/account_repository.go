package repositories

import "github.com/pablopasquim/CofrinhoPay/internal/core/domain"

type AccountRepository interface {
	Create(account *domain.Account) error
	GetByID(id uint) (*domain.Account, error)
	GetAll() ([]*domain.Account, error)
	Update(account *domain.Account) error
	Delete(id uint) error
}
