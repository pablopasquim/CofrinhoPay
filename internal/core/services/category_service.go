package services

import (
	"github.com/pablopasquim/CofrinhoPay/internal/core/domain"

	"github.com/pablopasquim/CofrinhoPay/internal/core/ports/repositories"
)

type CategoryService struct {
	categoryRepo repositories.CategoryRepository
}

func (s *CategoryService) CreateCategory(category *domain.Category) error {
	return s.categoryRepo.Create(category)
}

func (s *CategoryService) GetCategoryByID(id uint) (*domain.Category, error) {
	return s.categoryRepo.GetByID(id)
}

func (s *CategoryService) GetAllCategories() ([]*domain.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *CategoryService) UpdateCategory(category *domain.Category) error {
	return s.categoryRepo.Update(category)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.Delete(id)
}
