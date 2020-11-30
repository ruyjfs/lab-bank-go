package services

import (
	"github.com/ruyjfs/lab-bank-go/src/config"
	"github.com/ruyjfs/lab-bank-go/src/core"
	"github.com/ruyjfs/lab-bank-go/src/models"
	"gorm.io/gorm"
)

type Account struct {
	core.Service
}

func (a *Account) Find() (*gorm.DB, []models.Account) {
	var accounts []models.Account
	result := config.Db().Find(&accounts)
	return result, accounts
}
