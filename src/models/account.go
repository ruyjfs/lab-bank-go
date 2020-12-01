package models

import "github.com/ruyjfs/lab-bank-go/src/core"

type Account struct {
	DocumentNumber       uint          `json:"documentNumber" binding:"required" gorm:"unique;not null"`
	AvailableCreditLimit float64       `json:"availableCreditLimit" binding:"required"`
	Transactions         []Transaction `json:"transactions" binding:"required"`
	core.Model
}
