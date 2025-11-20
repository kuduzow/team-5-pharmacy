package models

import "time"

type Payment struct {
	Amount int 		`json:"amount"`
	Status string 	`json:"status"`
	Method string	`json:"method"`
	Paid_at time.Time `json:"paid_at"`
}

type CreatePayment struct {
	Amount int 		`json:"amount" binding:"required"`
	Status string 	`json:"status" binding:"required"` 
	Method string	`json:"method" binding:"required"`
	Paid_at time.Time `json:"paid_at" binding:"required"`
}


