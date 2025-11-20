package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FullName        string `json:"fullname"`
	Email           string `json:"email"`
	Phone           int    `json:"phone"`
	Default_address string `json:"default_address"`
}

type CreateUser struct {
	FullName        string `json:"fullname"`
	Email           string `json:"email"`
	Phone           int    `json:"phone"`
	Default_address string `json:"default_address"`
}
