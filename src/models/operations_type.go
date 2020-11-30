package models

import "github.com/ruyjfs/lab-bank-go/src/core"

type OperationsType struct {
	Description  string        `json:"description" binding:"required" gorm:"index"`
	Transactions []Transaction `json:"transactions" binding:"required" gorm:"index"`
	core.Model
}
