package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main(r *gin.Engine) *gin.Engine {

	r.POST("/account", func(c *gin.Context) {
		c.String(http.StatusOK, "AAAAAA")
	})

	r.GET("/accounts/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.String(http.StatusOK, "AAAAAA1", id)
	})

	r.POST("/transactions", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.String(http.StatusOK, "AAAAAA", id)
	})

	return r
}
