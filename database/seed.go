package database

import (
	"log"

	"github.com/ruyjfs/lab-bank-go/database/seeds"
)

func Seeder() {
	log.Printf("Seed: Start")
	seeds.OperationsTypes()
	log.Printf("Seed: Success")
}

// type Seed struct {
// 	Name string
// 	Run  func(*gorm.DB) error
// }

// func Seeder() []Seed {
// 	log.Printf("Seed: Start")
// 	return []Seed{
// 		{
// 			Name: "Operations Types"
// 			Run: seeds.OperationsTypes()
// 		}
// 	}
// }
