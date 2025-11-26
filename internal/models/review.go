package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID     uint     `json:"user_id"`
	User       User     `gorm:"foreignKey:UserID" json:"-"`
	MedicineID uint     `json:"medicine_id"`
	Medicine   Medicine `gorm:"foreignKey:MedicineID" json:"-"`
	Rating     int      `json:"rating"`
	Text       string   `json:"text"`
}

type ReviewCreateInput struct {
	UserID     uint   `json:"user_id"`
	MedicineID uint   `json:"medicine_id"`
	Rating     int    `json:"rating"`
	Text       string `json:"text"`
}

type ReviewUpdateInput struct {
	UserID     *uint   `json:"user_id"`
	MedicineID *uint   `json:"medicine_id"`
	Rating     *int    `json:"rating"`
	Text       *string `json:"text"`
}
