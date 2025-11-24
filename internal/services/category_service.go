package services

import (
	"errors"

	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrCategorytNotFound = errors.New("категория не найдена")
)

type CategoryService interface {
	CreateCategory(req models.CreateCategoryRequest) (*models.Category, error)
	UpdateCategory(id uint, req *models.UpdateCategoryRequest) (*models.Category, error)
	 DeleteCategory(id uint) error
	 GetByID(id uint) (*models.Category,error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (r *categoryService) CreateCategory(req models.CreateCategoryRequest) (*models.Category, error) {
	if err := r.validateCategoryCreate(&req); err != nil {
		return nil, err
	}

	category := models.Category{Name: req.Name}

	if err := r.repo.CreateCategory(&category); err != nil {
		return nil, err
	}
	return &category, nil

}

func (r *categoryService) validateCategoryCreate(req *models.CreateCategoryRequest) error {
	if req.Name == "" {
		return errors.New("название не может быть пустым")
	}
	return nil
}

func (r *categoryService) UpdateCategory(id uint, req *models.UpdateCategoryRequest) (*models.Category, error) {
	if err := r.applyCategoryValidate(req); err != nil {
		return nil, err
	}
	category, err := r.repo.GetById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategorytNotFound
		}
		return nil, err
	}

	if req.Name != nil {
		category.Name = *req.Name
	}

	if err := r.repo.UpdateCategory(category); err != nil {
		return nil, err
	}
	return category, nil

}

func (r *categoryService) applyCategoryValidate(req *models.UpdateCategoryRequest) error {
	if req.Name != nil {
		if *req.Name == "" {
			return errors.New("название не может быть пустым")
		}
	}
	return nil
}

func (r *categoryService) DeleteCategory(id uint) error {
	if _, err := r.repo.GetById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCategorytNotFound
		}
		return err
	}

	return r.repo.Delete(id)
}

func(r *categoryService) GetByID(id uint) (*models.Category,error){
	category, err := r.repo.GetById(id)

	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,ErrCategorytNotFound
		}
		return nil,err
	}
	return category,nil
}
