package repository

import (
	"errors"

	"github.com/kuduzow/team-5-pharmacy/internal/config/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetById(id uint) (*models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id uint) error
}

type gormPaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &gormPaymentRepository{db: db}
}

func (r *gormPaymentRepository) Create(payment *models.Payment) error {
	if payment == nil {
		return errors.New("payment отсутствует")
	}

	return r.db.Create(payment).Error
}

func (r *gormPaymentRepository) GetById(id uint) (*models.Payment, error) {
	var payment models.Payment

	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *gormPaymentRepository) Update(payment *models.Payment) error {
	if payment == nil {
		return nil
	}
	return r.db.Save(payment).Error
}

func (r *gormPaymentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payment{}, id).Error
}
