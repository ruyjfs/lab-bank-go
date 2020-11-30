package database

import (
	"log"

	"github.com/ruyjfs/lab-bank-go/src/database/seeds"
)

func Seeder() {
	log.Printf("Seed: Start")
	seeds.OperationsTypes()
	log.Printf("Seed: Success")
}
