package database

import (
	"log"

	"github.com/ruyjfs/lab-bank-go/src/config"
	"github.com/ruyjfs/lab-bank-go/src/models"
)

func Migrate() {
	log.Printf("Migrate: Start")
	db := config.Db()
	db.AutoMigrate(
		&models.Account{},
		&models.Transaction{},
		&models.OperationsType{},
	)
	log.Printf("Migrate: Success")
}
