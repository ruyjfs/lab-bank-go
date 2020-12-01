package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/lab-bank-go/src/config"
	"github.com/ruyjfs/lab-bank-go/src/core"
	"github.com/ruyjfs/lab-bank-go/src/models"
)

type Transaction struct {
	core.Controller
}

func (a *Transaction) Index(c *gin.Context) {
	var transactions []models.Transaction
	result := config.Db().Joins("Account").Joins("OperationsType").Find(&transactions)
	a.Json(c, result, transactions)
}

func (a *Transaction) Store(c *gin.Context) {
	var transaction *models.Transaction
	c.BindJSON(&transaction)
	if transaction.OperationsTypeID < 4 {

		var account models.Account
		config.Db().First(&account, transaction.AccountID)

		if transaction.Amount > account.AvailableCreditLimit {
			a.JsonFail(c, http.StatusBadRequest, "Ultrapassou o limite!")
			return
		}

		total := account.AvailableCreditLimit - transaction.Amount
		if total < 0 {
			total = float64(0)
		}

		account.AvailableCreditLimit = total
		config.Db().Save(&account)
		transaction.Amount = 0 - float64(transaction.Amount)
	}

	a.Json(c, config.Db().Create(&transaction), transaction)
}

func (a *Transaction) Update(c *gin.Context) {
	var transaction models.Transaction
	c.BindJSON(&transaction)
	a.Json(c, config.Db().Save(&transaction), transaction)
}

func (a *Transaction) Destroy(c *gin.Context) {
	var transaction models.Transaction
	config.Db().First(&transaction, c.Params.ByName("id"))
	a.Json(c, config.Db().Delete(&transaction), transaction)
}

func (a *Transaction) Show(c *gin.Context) {
	var transaction models.Transaction
	result := config.Db().First(&transaction, c.Params.ByName("id"))
	a.Json(c, result, transaction)
}
