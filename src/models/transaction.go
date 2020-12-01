package models

import (
	"time"

	"github.com/ruyjfs/lab-bank-go/src/core"
)

type Transaction struct {
	AccountID        int            `json:"accountId" binding:"required" gorm:"index;not null"`
	OperationsTypeID int            `json:"operationsTypeId" binding:"required" gorm:"not null"`
	Amount           float64        `json:"amount" binding:"required" gorm:"not null"`
	EventDate        *time.Time     `json:"eventDate" binding:"required"`
	Account          Account        `json:"account" binding:"required"`
	OperationsType   OperationsType `json:"operationsType" binding:"required"`
	core.Model
}
