package services

import (
	"fmt"

	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"github.com/pablopasquim/CofrinhoPay/internal/core/ports/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func (s *UserService) CreateUser(user *domain.User) error {

	if user.Username == "" {
		fmt.Println("Nome de usuário em branco")
	}

	return s.userRepo.Create(user)

}

func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
