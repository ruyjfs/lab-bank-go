package models

import (
	"time"

	"github.com/ruyjfs/lab-bank-go/core"
)

type Transaction struct {
	AccountID        int            `json:"accountId" binding:"required"`
	OperationsTypeID int            `json:"operationsTypeId" binding:"required"`
	Amount           float64        `json:"amount" binding:"required"`
	EventDate        *time.Time     `json:"eventDate" binding:"required"`
	Account          Account        `json:"account" binding:"required"`
	OperationsType   OperationsType `json:"operationsType" binding:"required"`
	core.Model
}
