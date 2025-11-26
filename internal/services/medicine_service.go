package services

import (
	"errors"

	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrMedicinetNotFound = errors.New("лекарство не найдено")
)
type MedicineService interface {
	CreateMedicine(req models.CreateMedicineRequest) (*models.Medicine, error)
	 UpdateMedicine(id uint, req *models.UpdateMedicineRequest) (*models.Medicine,error)
	GetMedecinesById(id uint) (*models.Medicine, error)
	 DeleteMedecineById(id uint )error
	 List(filter repository.MedicinesFilter) ([]models.Medicine, error)
	 
}
type medicineService struct {
	repo repository.MedicineRepository
}

func NewMedicineService(medecineRepo repository.MedicineRepository) MedicineService {
	return &medicineService{repo: medecineRepo}
}

func (m *medicineService) CreateMedicine(req models.CreateMedicineRequest) (*models.Medicine, error) {
if err:= m.validateMedicineCreate(&req);err!=nil{
	return nil,err
}
	medicine := models.Medicine{
		Name:                 req.Name,
        Description:          req.Description,
        Price:                req.Price,
        InStock:              req.InStock,
        StockQuantity:        req.StockQuantity,
        CategoryId:           req.CategoryId,    
        SubcategoryId:        req.SubcategoryId, 
        Manufacturer:         req.Manufacturer,
        PrescriptionRequired: req.PrescriptionRequired,
	}

	if err := m.repo.Create(&medicine); err != nil {
		return nil, err

	}
	return &medicine, nil
}
func (m *medicineService) validateMedicineCreate(req *models.CreateMedicineRequest) error {
	if req.Name == "" {
		return  errors.New("название лекарства не может быть пустым")
	}
	if req.Price <= 0 {
		return  errors.New("цена должна быть положительной")
	}
	if req.StockQuantity < 0 {
		return  errors.New("количество не должно быть отрицательным")
	}
	return nil
}

func (m *medicineService) UpdateMedicine(id uint, req *models.UpdateMedicineRequest) (*models.Medicine, error) {
    if req == nil {
        return m.GetMedecinesById(id)
    }
    if err := m.applyMedicinesValidate(req); err != nil {
        return nil, err
    }

    cur, err := m.repo.GetMedecinesById(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrMedicinetNotFound
        }
        return nil, err
    }

    if req.Name != nil { cur.Name = *req.Name }
    if req.Description != nil { cur.Description = *req.Description }
    if req.Price != nil { cur.Price = *req.Price }
    if req.InStock != nil { cur.InStock = *req.InStock }                  
    if req.StockQuantity != nil { cur.StockQuantity = *req.StockQuantity }
    if req.CategoryId != nil { cur.CategoryId = *req.CategoryId }
    if req.SubcategoryId != nil { cur.SubcategoryId = *req.SubcategoryId }
    if req.Manufacturer != nil { cur.Manufacturer = *req.Manufacturer }
    if req.PrescriptionRequired != nil { cur.PrescriptionRequired = *req.PrescriptionRequired }

    if err := m.repo.UpdateCategoryRequest(cur); err != nil {
        return nil, err
    }
    return cur, nil
}

func (m *medicineService) applyMedicinesValidate(req *models.UpdateMedicineRequest) error {
    if req.Name != nil && *req.Name == "" {
        return errors.New("название лекарства не может быть пустым")
    }
    if req.Price != nil && *req.Price <= 0 {
        return errors.New("цена должна быть положительной")
    }
   
    if req.StockQuantity != nil && *req.StockQuantity < 0 {
        return errors.New("количество не должно быть отрицательным")
    }
    return nil
}

func (m *medicineService) GetMedecinesById(id uint) (*models.Medicine, error){
 medicine, err:= m.repo.GetMedecinesById(id)

 if err!=nil{
	if errors.Is(err,gorm.ErrRecordNotFound){
	return nil, ErrMedicinetNotFound
 }
 return nil, err
 }
 return medicine,nil

}


func (m *medicineService) DeleteMedecineById(id uint) error {
	if _, err := m.repo.GetMedecinesById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrMedicinetNotFound
		}
		return err
	}
	if err := m.repo.DeleteMedecines(id); err != nil {
		return err
	}
	return nil
}

func (m *medicineService) List(filter repository.MedicinesFilter) ([]models.Medicine, error) {
	return m.repo.List(filter)
}
