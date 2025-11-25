package repository

import (
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"gorm.io/gorm"
)

type SubcategoryRepository interface {
	Create(subcat *models.Subcategory) error
	GetByID(id uint) (*models.Subcategory, error)
	Update(subcat *models.Subcategory) error
	Delete(id uint) error
	ListByCategoryID(categoryID uint) ([]models.Subcategory, error)
}

type gormSubcategoryRepository struct {
	db *gorm.DB
}

func NewSubcategoryRepository(db *gorm.DB) SubcategoryRepository {
	return &gormSubcategoryRepository{db: db}
}

func (r *gormSubcategoryRepository) Create(subcat *models.Subcategory) error {
	return r.db.Create(subcat).Error
}

func (r *gormSubcategoryRepository) GetByID(id uint) (*models.Subcategory, error) {
	var subcat models.Subcategory
	if err := r.db.First(&subcat, id).Error; err != nil {
		return nil, err
	}
	return &subcat, nil
}

func (r *gormSubcategoryRepository) Update(subcat *models.Subcategory) error {
	return r.db.Save(subcat).Error
}

func (r *gormSubcategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Subcategory{}, id).Error
}

func (r *gormSubcategoryRepository) ListByCategoryID(categoryID uint) ([]models.Subcategory, error) {
    var subs []models.Subcategory
    if err := r.db.Where("category_id = ?", categoryID).Find(&subs).Error; err != nil {
        return nil, err
    }
    return subs, nil
}
