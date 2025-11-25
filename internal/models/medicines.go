package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Name                 string      `json:"name"`
	Description          string      `json:"description"`
	Price                int         `json:"price"`
	InStock              bool        `json:"instock"`
	StockQuantity        uint        `json:"stockquantity"`
	CategoryId           uint        `json:"categoryid"`
	Category             Category    `gorm:"foreignKey:CategoryId" json:"-"`
	SubcategoryId        uint        `json:"subcategoryid"`
	Subcategory          Subcategory `gorm:"foreignKey:SubcategoryId" json:"-"`
	Manufacturer         string      `json:"manufacturer"`
	PrescriptionRequired bool        `json:"prescriptionrequired"`
	AvgRating            float64     `gorm:"-" json:"avgraiting"`
}

type CreateMedicineRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                int     `json:"price"`
	InStock              bool    `json:"instock"`
	StockQuantity        uint    `json:"stockquantity"`
	CategoryId           uint    `json:"categoryid"`
	SubcategoryId        uint    `json:"subcategoryid"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescriptionrequired"`
	AvgRating            float64 `gorm:"-" json:"avgraiting"`
}

type UpdateMedicineRequest struct {
	Name                 *string  `json:"name"`
	Description          *string  `json:"description"`
	Price                *int     `json:"price"`
	InStock              *bool    `json:"instock"`
	StockQuantity        *uint    `json:"stockquantity"`
	CategoryId           *uint    `json:"categoryid"`
	SubcategoryId        *uint    `json:"subcategoryid"`
	Manufacturer         *string  `json:"manufacturer"`
	PrescriptionRequired *bool    `json:"prescriptionrequired"`
	AvgRating            *float64 `gorm:"-" json:"avgraiting"`
}
