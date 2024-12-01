package service

import (
	"personal-finance-app/internal/models"
	"personal-finance-app/internal/repository"
)

type CategoryService interface {
	GetCategories(userID uint) ([]*models.Category, error)
	CreateCategory(userID uint, name string) error
	UpdateCategory(categoryID uint, name string) error
	DeleteCategory(categoryID uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetCategories(userID uint) ([]*models.Category, error) {
	return s.repo.GetCategoriesByUser(userID)
}

func (s *categoryService) CreateCategory(userID uint, name string) error {
	category := &models.Category{UserID: userID, Name: name}
	return s.repo.CreateCategory(category)
}

func (s *categoryService) UpdateCategory(categoryID uint, name string) error {
	category, err := s.repo.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}
	category.Name = name
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(categoryID uint) error {
	return s.repo.DeleteCategory(categoryID)
}
