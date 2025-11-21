package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	Price                 int     `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        uint    `json:"stock_quantity"`
	CategoryId           string  `json:"category_id"`
	SubcategoryId        string  `json:"subcategory_id"`
	Manufacturer          string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `gorm:"-" json:"avg_raiting"`
}

type CreateMedicineRequest struct {
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	Price                 int     `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        uint    `json:"stock_quantity"`
	CategoryId           string  `json:"category_id"`
	SubcategoryId        string  `json:"subcategory_id"`
	Manufacturer          string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `gorm:"-" json:"avg_raiting"`
}
type UpdateMedicineRequest struct {
	Name                  *string  `json:"name"`
	Description           *string  `json:"description"`
	Price                 *int     `json:"price"`
	InStock              *bool    `json:"in_stock"`
	StockQuantity        *uint    `json:"stock_quantity"`
	CategoryId           *string  `json:"category_id"`
	SubcategoryId        *string  `json:"subcategory_id"`
	Manufacturer          *string  `json:"manufacturer"`
	PrescriptionRequired *bool    `json:"prescription_required"`
	AvgRating            *float64 `gorm:"-" json:"avg_raiting"`
}
