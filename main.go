package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/lab-bank-go/routes"
)

func main() {
	routes.Graphql(
		routes.Main(
			gin.Default(),
		),
	).Run()
}
