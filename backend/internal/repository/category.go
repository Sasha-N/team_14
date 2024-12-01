package repository

import (
	"personal-finance-app/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUser(userID uint) ([]*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(categoryID uint) error
	GetCategoryByID(categoryID uint) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetCategoriesByUser(userID uint) ([]*models.Category, error) {
	var categories []*models.Category
	err := r.db.Where("user_id = ?", userID).Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(categoryID uint) error {
	return r.db.Delete(&models.Category{}, categoryID).Error
}

func (r *categoryRepository) GetCategoryByID(categoryID uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, categoryID).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
