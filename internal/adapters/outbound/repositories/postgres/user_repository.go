package repository

import (
	"errors"

	"github.com/pablopasquim/CofrinhoPay/internal/adapters/outbound/repositories/models"
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
	userDB := models.UserFromDomain(user)
	if err := r.db.Create(userDB).Error; err != nil {
		return err
	}

	user.ID = userDB.ID
	user.CreatedAt = userDB.CreatedAt
	user.UpdatedAt = userDB.UpdatedAt
	return nil
}

func (r *UserRepositoryPostgres) GetByID(id uint) (*domain.User, error) {
	var userDB models.UserDB
	if err := r.db.First(&userDB, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return userDB.ToDomain(), nil
}

func (r *UserRepositoryPostgres) GetByUsername(username string) (*domain.User, error) {
	var userDB models.UserDB
	if err := r.db.Where("username = ?", username).First(&userDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return userDB.ToDomain(), nil
}

func (r *UserRepositoryPostgres) GetByEmail(email string) (*domain.User, error) {
	var userDB models.UserDB
	if err := r.db.Where("email = ?", email).First(&userDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return userDB.ToDomain(), nil
}

func (r *UserRepositoryPostgres) GetAll() ([]*domain.User, error) {
	var usersDB []models.UserDB
	if err := r.db.Find(&usersDB).Error; err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(usersDB))
	for i, userDB := range usersDB {
		users[i] = userDB.ToDomain()
	}
	return users, nil
}

func (r *UserRepositoryPostgres) Update(user *domain.User) error {
	userDB := models.UserFromDomain(user)
	if err := r.db.Save(userDB).Error; err != nil {
		return err
	}
	// Atualiza os timestamps do domain user
	user.UpdatedAt = userDB.UpdatedAt
	return nil
}

func (r *UserRepositoryPostgres) Delete(id uint) error {
	if err := r.db.Delete(&models.UserDB{}, id).Error; err != nil {
		return err
	}
	return nil
}
