package services

import (
	"errors"

	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"github.com/pablopasquim/CofrinhoPay/internal/core/ports/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}

	existingUser, _ := s.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email já está em uso")
	}

	return s.userRepo.Create(user)
}

func (s *UserService) validateUser(user *domain.User) error {
	if user.Email == "" {
		return errors.New("email é obrigatório")
	}

	if user.Username == "" {
		return errors.New("username é obrigatório")
	}

	if user.Password == "" {
		return errors.New("senha é obrigatória")
	}

	if len(user.Password) < 6 {
		return errors.New("senha deve ter pelo menos 6 caracteres")
	}

	return nil
}

func (s *UserService) Login(email, password string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("senha incorreta")
	}

	return user, nil
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
