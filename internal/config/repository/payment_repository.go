package repository

import (
	"github.com/kuduzow/team-5-pharmacy/internal/config/models"
	"gorm.io/gorm"
)

type PaymentRepository interface{

}

type paymentRepository struct{
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository{
	return paymentRepository{db}
}

func (r *paymentRepository) Create(payment *models.Payment) error{
	if payment == nil{
		return nil
	}
	
	return r.db.Create(payment).Error
}

func(r *paymentRepository) GetById (id uint) (*models.Payment,error){
	var payment models.Payment

	if err := r.db.First(&payment,id).Error; err != nil{
		return nil, err
	}
	return &payment,nil
}

func(r *paymentRepository) Update(payment *models.Payment) error{
	if payment == nil{
		return nil
	}
	return r.db.Save(payment).Error
}

	