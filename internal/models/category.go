package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name *string `json:"name"`
}
