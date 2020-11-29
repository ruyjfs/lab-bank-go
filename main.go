package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/lab-bank-go/database"
	"github.com/ruyjfs/lab-bank-go/routes"
)

func main() {

	database.Migrate()
	database.Seeder()

	routes.Graphql(
		routes.Main(
			gin.Default(),
		),
	).Run()
}
