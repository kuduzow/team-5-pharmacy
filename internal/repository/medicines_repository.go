package repository

import (
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"gorm.io/gorm"
)

type MedicinesFilter struct {
	CategoryId    *uint
	SubcategoryId *uint
	InStock       *bool
}
type MedicineRepository interface {
	Create(medicine *models.Medicine) error
	UpdateCategoryRequest(medicine *models.Medicine) error
	DeleteMedecines(id uint) error
	GetMedecinesById(id uint) (*models.Medicine, error)
	List(filter MedicinesFilter) ([]models.Medicine, error)
	Exists(id uint) (bool, error)
}

type gormMedicineRepository struct {
	db *gorm.DB
}

func NewMedicinesRepository(db *gorm.DB) MedicineRepository {
	return &gormMedicineRepository{db: db}
}

func (r *gormMedicineRepository) Create(medicine *models.Medicine) error {
	if medicine == nil {

		return nil
	}
	return r.db.Create(medicine).Error
}
func (r *gormMedicineRepository) UpdateCategoryRequest(medicine *models.Medicine) error {
	if medicine == nil {

		return nil
	}
	return r.db.Save(medicine).Error
}
func (r *gormMedicineRepository) DeleteMedecines(id uint) error {
	if err := r.db.Delete(&models.Medicine{}, id).Error; err != nil {
		return err

	}
	return nil
}

func (r *gormMedicineRepository) GetMedecinesById(id uint) (*models.Medicine, error) {
	var med models.Medicine
	if err := r.db.First(&med, id).Error; err != nil {
		return nil, err
	}
	return &med, nil

}
func (r *gormMedicineRepository) Exists(id uint) (bool, error) {
	var count int64
	err := r.db.
		Model(&models.Medicine{}).
		Where("id = ?", id).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *gormMedicineRepository) List(filter MedicinesFilter) ([]models.Medicine, error) {
	query := r.db.Model(&models.Medicine{})

	var medicines []models.Medicine

	if filter.CategoryId != nil {
		query = query.Where("category_id = ?", *filter.CategoryId)
	}

	if filter.SubcategoryId != nil {
		query = query.Where("subcategory_id = ?", *filter.SubcategoryId)
	}

	if filter.InStock != nil {
		query = query.Where("in_stock = ?", *filter.InStock)
	}

	if err := query.Find(&medicines).Error; err != nil {
		return nil, err
	}
	return medicines, nil
}
