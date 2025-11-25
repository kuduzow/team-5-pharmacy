package models

import "gorm.io/gorm"

type Subcategory struct {
	gorm.Model
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

type CreateSubcategoryRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type UpdateSubcategoryRequest struct {
	Name *string `json:"name"`
}
