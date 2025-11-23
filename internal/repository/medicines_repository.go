package repository

import (
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	Create(medicine *models.Medicine) error
	UpdateCategoryRequest(medicine *models.Medicine) error
	DeleteMedecines(medicine *models.Medicine) error
	GetMedecines(medicine *models.Medicine) error
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
func (r *gormMedicineRepository) DeleteMedecines(medicine *models.Medicine) error {
	if medicine == nil {

		return nil
	}
	return r.db.Delete(medicine).Error
}
func (r *gormMedicineRepository) GetMedecines(medicine *models.Medicine) error {
	if medicine == nil {

		return nil
	}
	return r.db.First(medicine).Error
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
