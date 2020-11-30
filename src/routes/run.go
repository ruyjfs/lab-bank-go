package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"graphql-playground": "http://localhost:8085/graphiql",
			"graphql":            "http://localhost:8085/graphql",
			"v1":                 "http://localhost:8085/api/v1",
		})
	})

	V1(r)
	Graphql(r)
	r.Run()
}
