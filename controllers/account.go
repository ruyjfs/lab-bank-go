package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/lab-bank-go/config"
	"github.com/ruyjfs/lab-bank-go/core"
	"github.com/ruyjfs/lab-bank-go/models"
)

type Account struct {
	core.Controller
}

func (a *Account) Index(c *gin.Context) {
	var accounts []models.Account
	result := config.Db().Find(&accounts)
	a.Json(c, result, accounts)
}

func (a *Account) Store(c *gin.Context) {
	var account *models.Account
	c.BindJSON(&account)
	a.Json(c, config.Db().Create(&account), account)
}

func (a *Account) Update(c *gin.Context) {
	var account models.Account
	c.BindJSON(&account)
	a.Json(c, config.Db().Save(&account), account)
}

func (a *Account) Destroy(c *gin.Context) {
	var account models.Account
	config.Db().First(&account, c.Params.ByName("id"))
	a.Json(c, config.Db().Delete(&account), account)
}

func (a *Account) Show(c *gin.Context) {
	var account models.Account
	result := config.Db().First(&account, c.Params.ByName("id"))
	a.Json(c, result, account)
}
