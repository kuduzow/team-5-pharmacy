package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	Price                 int     `json:"price"`
	In_stock              bool    `json:"in_stock"`
	Stock_quantity        uint    `json:"stock_quantity"`
	Category_id           string  `json:"category_id"`
	Subcategory_id        string  `json:"subcategory_id"`
	Manufacturer          string  `json:"manufacturer"`
	Prescription_required bool    `json:"prescription_required"`
	Avg_rating            float64 `gorm:"-" json:"avg_raiting"`
}

type CreateMedicineRequest struct {
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	Price                 int     `json:"price"`
	In_stock              bool    `json:"in_stock"`
	Stock_quantity        uint    `json:"stock_quantity"`
	Category_id           string  `json:"category_id"`
	Subcategory_id        string  `json:"subcategory_id"`
	Manufacturer          string  `json:"manufacturer"`
	Prescription_required bool    `json:"prescription_required"`
	Avg_rating            float64 `gorm:"-" json:"avg_raiting"`
}
type UpdateMedicineRequest struct {
	Name                  *string  `json:"name"`
	Description           *string  `json:"description"`
	Price                 *int     `json:"price"`
	In_stock              *bool    `json:"in_stock"`
	Stock_quantity        *uint    `json:"stock_quantity"`
	Category_id           *string  `json:"category_id"`
	Subcategory_id        *string  `json:"subcategory_id"`
	Manufacturer          *string  `json:"manufacturer"`
	Prescription_required *bool    `json:"prescription_required"`
	Avg_rating            *float64 `gorm:"-" json:"avg_raiting"`
}
