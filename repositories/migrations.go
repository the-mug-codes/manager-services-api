package repositories

import (
	"log"
	"time"

	security "github.com/the-mug-codes/adapters-service-api/security"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	"gorm.io/gorm"
)

func Migrations(database *gorm.DB) {
	err := database.AutoMigrate(
		&security.User{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.Label{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.Person{},
		&entity.PersonExtra{},
		&entity.PersonAddress{},
		&entity.PersonEmail{},
		&entity.PersonPhone{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.Product{},
		&entity.ProductPrice{},
		&entity.ProductCategory{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.Project{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.Contract{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.HelpDeskTicket{},
		&entity.HelpDeskTicketMessage{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.Invoice{},
		&entity.InvoiceItem{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(
		&entity.ChatSession{},
		&entity.ChatMessage{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("\033[person migrate finished at %s\033[0m", time.Now().UTC())
}
