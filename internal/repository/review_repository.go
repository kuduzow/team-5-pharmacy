package repository

import (
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *models.Review) error

	GetByID(id uint) (*models.Review, error)

	Update(review *models.Review) error

	Delete(id uint) error
}

type gormReviewRepository struct {
	db *gorm.DB
}

func (r *gormReviewRepository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *gormReviewRepository) GetByID(id uint) (*models.Review, error) {
	var review models.Review
	if err := r.db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *gormReviewRepository) Update(review *models.Review) error {
	return r.db.Save(review).Error
}

func (r *gormReviewRepository) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}
