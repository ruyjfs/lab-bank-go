package seeds

import (
	"log"

	"github.com/ruyjfs/lab-bank-go/src/config"
	"github.com/ruyjfs/lab-bank-go/src/models"
)

func OperationsTypes() {
	db := config.Db()
	var operationsTypes []models.OperationsType

	result := db.Find(&operationsTypes)
	if result.RowsAffected >= 4 {
		log.Printf("Seed (OperationsTypes): Nothing to seed")
		return
	}

	operationsTypes = []models.OperationsType{
		{Description: "COMPRA A VISTA"},
		{Description: "COMPRA PARCELADA"},
		{Description: "SAQUE"},
		{Description: "PAGAMENTO"},
	}

	db.Create(&operationsTypes)
	log.Printf("Seed (OperationsTypes): Success")
}
