package services

import (
	"errors"

	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrSubcategorytNotFound = errors.New("субкатегория не найдена")
)

type SubcategoryService interface {
	CreateSubcategory(req models.CreateSubcategoryRequest) (*models.Subcategory, error)
	GetSubcategoryByID(id uint) (*models.Subcategory, error)
	UpdateSubcategory(id uint, req models.UpdateSubcategoryRequest) (*models.Subcategory, error)
	DeleteSubcategory(id uint) error
	ListSubcategoriesByCategory(categoryID uint) ([]models.Subcategory, error)
}

type subcategoryService struct {
	repo repository.SubcategoryRepository
}

func NewSubcategoryService(repo repository.SubcategoryRepository) SubcategoryService {
	return &subcategoryService{repo: repo}
}

func (s *subcategoryService) CreateSubcategory(req models.CreateSubcategoryRequest) (*models.Subcategory, error) {
	subcat := &models.Subcategory{
		Name:       req.Name,
		CategoryID: req.CategoryID,
	}
	if err := s.repo.Create(subcat); err != nil {
		return nil, err
	}
	return subcat, nil
}

func (s *subcategoryService) GetSubcategoryByID(id uint) (*models.Subcategory, error) {
	return s.repo.GetByID(id)
}

func (s *subcategoryService) UpdateSubcategory(id uint, req models.UpdateSubcategoryRequest) (*models.Subcategory, error) {
	subcat, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil, ErrSubcategorytNotFound
		}
		return nil, err
	}
	if req.Name != nil {
		subcat.Name = *req.Name
	}
	if err := s.repo.Update(subcat); err != nil {
		return nil, err
	}
	return subcat, nil
}

func (s *subcategoryService) DeleteSubcategory(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return ErrSubcategorytNotFound
		}
		return err
	}
	return s.repo.Delete(id)
}

func (s *subcategoryService) ListSubcategoriesByCategory(categoryID uint) ([]models.Subcategory, error) {
	if _, err := s.repo.GetByID(categoryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategorytNotFound
		}
		return nil, err
	}
	return s.repo.ListByCategoryID(categoryID)
}
