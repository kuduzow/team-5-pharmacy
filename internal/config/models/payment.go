package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
)

type PaymentMethod string

const (
	PaymentMethodCard         PaymentMethod = "card"
	PaymentMethodCash         PaymentMethod = "cash"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

type Payment struct {
	gorm.Model
	Amount  int           `json:"amount"`
	Status  PaymentStatus `json:"-"`
	Method  PaymentMethod `json:"-"`
	Paid_at time.Time     `json:"paid_at"`
}

type CreatePaymentRequest struct {
	Amount  int       `json:"amount" binding:"required"`
	Status  PaymentStatus    `json:"status" binding:"required"`
	Method  PaymentMethod    `json:"method" binding:"required"`
	Paid_at time.Time `json:"paid_at" binding:"required"`
}

type UpdatePaymentRequest struct {
	Amount  *int       `json:"amount"`
	Status  *PaymentStatus    `json:"status"`
	Method  *PaymentMethod    `json:"method"`
	Paid_at *time.Time `json:"paid_at"`
}
