package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/lab-bank-go/config"
	"github.com/ruyjfs/lab-bank-go/core"
	"github.com/ruyjfs/lab-bank-go/models"
)

type OperationType struct {
	core.Controller
}

func (a *OperationType) Index(c *gin.Context) {
	var operationsTypes []models.OperationsType
	result := config.Db().Find(&operationsTypes)
	a.Json(c, result, operationsTypes)
}

func (a *OperationType) Store(c *gin.Context) {
	var operationType *models.OperationsType
	c.BindJSON(&operationType)
	a.Json(c, config.Db().Create(&operationType), operationType)
}

func (a *OperationType) Update(c *gin.Context) {
	var operationType models.OperationsType
	c.BindJSON(&operationType)
	a.Json(c, config.Db().Save(&operationType), operationType)
}

func (a *OperationType) Destroy(c *gin.Context) {
	var operationType models.OperationsType
	config.Db().First(&operationType, c.Params.ByName("id"))
	a.Json(c, config.Db().Delete(&operationType), operationType)
}

func (a *OperationType) Show(c *gin.Context) {
	var operationType models.OperationsType
	result := config.Db().First(&operationType, c.Params.ByName("id"))
	a.Json(c, result, operationType)
}
