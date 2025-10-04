package services

import (
	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"github.com/pablopasquim/CofrinhoPay/internal/core/ports/repositories"
)

type AccountService struct {
	accountRepo repositories.AccountRepository
}

func (s *AccountService) CreateAccount(account *domain.Account) error {
	return s.accountRepo.Create(account)
}

func (s *AccountService) GetAccountByID(id uint) (*domain.Account, error) {
	return s.accountRepo.GetByID(id)
}

func (s *AccountService) GetAllAccounts() ([]*domain.Account, error) {
	return s.accountRepo.GetAll()
}

func (s *AccountService) UpdateAccount(account *domain.Account) error {
	return s.accountRepo.Update(account)
}

func (s *AccountService) DeleteAccount(id uint) error {
	return s.accountRepo.Delete(id)
}
