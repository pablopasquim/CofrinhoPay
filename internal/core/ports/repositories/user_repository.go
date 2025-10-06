package repositories

import "github.com/pablopasquim/CofrinhoPay/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id uint) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}
