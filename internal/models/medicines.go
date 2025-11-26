package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model

	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                int     `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int    `json:"stock_quantity"`
	CategoryId           uint    `json:"category_id"`
	SubcategoryId        uint    `json:"subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `gorm:"-" json:"avg_rating"`
}

type CreateMedicineRequest struct {
	Name                 string  `json:"name" binding:"required"`
	Description          string  `json:"description"`
	Price                int     `json:"price" binding:"required,gt=0"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int    `json:"stock_quantity"`
	CategoryId           uint    `json:"category_id" binding:"required"`
	SubcategoryId        uint    `json:"subcategory_id" binding:"required"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `gorm:"-" json:"avg_rating"`
}

type UpdateMedicineRequest struct {
	Name                 *string  `json:"name"`
	Description          *string  `json:"description"`
	Price                *int     `json:"price"`
	InStock              *bool    `json:"in_stock"`
	StockQuantity        *int    `json:"stock_quantity"`
	CategoryId           *uint    `json:"category_id"`
	SubcategoryId        *uint    `json:"subcategory_id"`
	Manufacturer         *string  `json:"manufacturer"`
	PrescriptionRequired *bool    `json:"prescription_required"`
	AvgRating            *float64 `gorm:"-" json:"avg_rating"`
}
