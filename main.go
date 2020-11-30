package main

import (
	"github.com/ruyjfs/lab-bank-go/src/database"
	"github.com/ruyjfs/lab-bank-go/src/routes"
)

func main() {

	database.Migrate()
	database.Seeder()

	routes.Run()
}
