package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/lab-bank-go/controllers"
)

func Main(r *gin.Engine) *gin.Engine {

	account := new(controllers.Account)
	r.GET("/accounts", account.Index)
	r.POST("/accounts", account.Store)
	r.PATCH("/accounts", account.Update)
	r.DELETE("/accounts/:id", account.Destroy)
	r.GET("/accounts/:id", account.Show)

	transaction := new(controllers.Transaction)
	r.GET("/transactions", transaction.Index)
	r.POST("/transactions", transaction.Store)
	r.PATCH("/transactions", transaction.Update)
	r.DELETE("/transactions/:id", transaction.Destroy)
	r.GET("/transactions/:id", transaction.Show)

	operationType := new(controllers.OperationType)
	r.GET("/operations-types", operationType.Index)
	r.POST("/operations-types", operationType.Store)
	r.PATCH("/operations-types", operationType.Update)
	r.DELETE("/operations-types/:id", operationType.Destroy)
	r.GET("/operations-types/:id", operationType.Show)

	return r
}
