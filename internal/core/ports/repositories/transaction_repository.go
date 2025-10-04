package repositories

import "github.com/pablopasquim/CofrinhoPay/internal/core/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	GetByID(id uint) (*domain.Transaction, error)
	GetAll() ([]*domain.Transaction, error)
	Update(transaction *domain.Transaction) error
	Delete(id uint) error
}
