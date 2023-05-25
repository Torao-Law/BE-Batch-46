package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategory(ID int) (models.Category, error)
	FindCategories() ([]models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.Preload("Product").First(&category, ID).Error

	return category, err
}

func (r *repository) FindCategories() ([]models.Category, error) {
	var category []models.Category
	err := r.db.Preload("Product").Find(&category).Error

	return category, err
}
