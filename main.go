package main

import (
	"github.com/ruyjfs/lab-bank-go/database"
	"github.com/ruyjfs/lab-bank-go/routes"
)

func main() {

	database.Migrate()
	database.Seeder()

	routes.Run()
}
