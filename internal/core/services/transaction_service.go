package services

import (
	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"github.com/pablopasquim/CofrinhoPay/internal/core/ports/repositories"
)

type TransactionService struct {
	transactionRepo repositories.TransactionRepository
}

func (s *TransactionService) CreateTransaction(transaction *domain.Transaction) error {
	return s.transactionRepo.Create(transaction)
}

func (s *TransactionService) GetTransactionByID(id uint) (*domain.Transaction, error) {
	return s.transactionRepo.GetByID(id)
}

func (s *TransactionService) GetAllTransactions() ([]*domain.Transaction, error) {
	return s.transactionRepo.GetAll()
}

func (s *TransactionService) UpdateTransaction(transaction *domain.Transaction) error {
	return s.transactionRepo.Update(transaction)
}

func (s *TransactionService) DeleteTransaction(id uint) error {
	return s.transactionRepo.Delete(id)
}
