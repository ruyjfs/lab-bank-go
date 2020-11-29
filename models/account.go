package models

import "github.com/ruyjfs/lab-bank-go/core"

type Account struct {
	DocumentNumber uint          `json:"documentNumber" binding:"required"`
	Transactions   []Transaction `json:"transactions" binding:"required"`
	core.Model
}
