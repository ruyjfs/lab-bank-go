package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/lab-bank-go/controllers"
)

func V1(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	{
		account := new(controllers.Account)
		v1.GET("/accounts", account.Index)
		v1.POST("/accounts", account.Store)
		v1.PATCH("/accounts", account.Update)
		v1.DELETE("/accounts/:id", account.Destroy)
		v1.GET("/accounts/:id", account.Show)

		transaction := new(controllers.Transaction)
		v1.GET("/transactions", transaction.Index)
		v1.POST("/transactions", transaction.Store)
		v1.PATCH("/transactions", transaction.Update)
		v1.DELETE("/transactions/:id", transaction.Destroy)
		v1.GET("/transactions/:id", transaction.Show)

		operationType := new(controllers.OperationType)
		v1.GET("/operations-types", operationType.Index)
		v1.POST("/operations-types", operationType.Store)
		v1.PATCH("/operations-types", operationType.Update)
		v1.DELETE("/operations-types/:id", operationType.Destroy)
		v1.GET("/operations-types/:id", operationType.Show)
	}

	return r
}
