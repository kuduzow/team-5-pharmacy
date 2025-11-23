package repository

import (
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface{
	CreateCategory(category *models.CreateCategoryRequest) error
	UpdateCategoryRequest(category *models.UpdateCategoryRequest) error
	GetById(id uint) (*models.Category,error)
	Delete(id uint) error
}

type gormCaregoryRepository struct{
	db *gorm.DB
}

func NewCategoryRepository (db *gorm.DB) CategoryRepository{
	return &gormCaregoryRepository{db:db}
}

func(r *gormCaregoryRepository) CreateCategory(category *models.CreateCategoryRequest) error{
	return r.db.Create(category).Error
}

func (r *gormCaregoryRepository) UpdateCategoryRequest(category *models.UpdateCategoryRequest) error{
	if category == nil{
		return nil
	}

	return r.db.Save(category).Error
}

func(r *gormCaregoryRepository) GetById(id uint) (*models.Category,error){
	var category *models.Category

	if err := r.db.First(&category,id).Error; err !=nil{
		return nil,err
	}
	return category,nil
}

func(r *gormCaregoryRepository) Delete(id uint) error{
	return r.db.Delete(&models.Category{},id).Error
}

