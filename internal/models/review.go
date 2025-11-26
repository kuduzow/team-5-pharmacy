package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID     string `json:"user_id"`
	MedicineID uint   `json:"medicine_id"`
	Rating     int    `json:"rating"`
	Text       string `json:"text"`
}

type ReviewCreateInput struct {
	UserID     string `json:"user_id"`
	MedicineID uint   `json:"medicine_id"`
	Rating     int    `json:"rating"`
	Text       string `json:"text"`
}

type ReviewUpdateInput struct {
	UserID     *string `json:"user_id"`
	MedicineID *uint   `json:"medicine_id"`
	Rating     *int    `json:"rating"`
	Text       *string `json:"text"`
}
