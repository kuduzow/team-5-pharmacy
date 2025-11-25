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
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		StockQuantity: req.StockQuantity,
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
	if req.StockQuantity <= 0 {
		return  errors.New("должно быть больше 0")
	}
	return nil
}
func (m *medicineService) UpdateMedicine(id uint ,req *models.UpdateMedicineRequest) (*models.Medicine, error) {

	if err:= m.applyMedicinesValidate(req);err!=nil{
		return nil,err
	}
	medicine := models.Medicine{
		Name:          *req.Name,
		Description:  *req.Description,
		Price:         *req.Price,
		StockQuantity: *req.StockQuantity,
	}
	

	if err := m.repo.UpdateCategoryRequest(&medicine); err != nil {
		return nil, err

	}
	return &medicine, nil
}
func (m *medicineService) applyMedicinesValidate(req *models.UpdateMedicineRequest) error {

	if req.Name !=nil && *req.Name == "" {
		return errors.New("название лекарства не может быть пустым")
	}
	if req.Price!=nil && *req.Price <=0 {
		return  errors.New("цена должна быть положительной")
	}
	if req.StockQuantity != nil && *req.StockQuantity==0 {
		return  errors.New("должно быть больше 0")
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
