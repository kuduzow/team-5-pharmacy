package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	DefaultAddress string `json:"default_address"`
}

type CreateUserRequest struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	DefaultAddress string `json:"default_address"`
}

type UpdateUserRequest struct {
	FullName       *string `json:"full_name"`
	Email          *string `json:"email"`
	Phone          *string `json:"phone"`
	DefaultAddress *string `json:"default_address"`
}
