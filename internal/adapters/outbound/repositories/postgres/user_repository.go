// adapters/postgres/repository/user_repository.go
package repository

import (
	"errors"

	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"
	"github.com/pablopasquim/CofrinhoPay/internal/core/ports/repositories"
	"gorm.io/gorm"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepositoryPostgres(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) Create(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryPostgres) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPostgres) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryPostgres) Update(user *domain.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryPostgres) Delete(id uint) error {
	if err := r.db.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryPostgres) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
