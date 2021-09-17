package db

import (
	"log"

	"github.com/OliveiraJ/ecoin-tracker-v2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creates a table of models.Read struct into the database
func Setup(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.Read{})
	if err != nil {
		panic(err)
	}

}

// Start a new connection with the database
func NewDbConnection() *gorm.DB {
	dsn := "host=172.21.0.2 user=ecointracker password=jJ99268940*/ecointracker dbname=ecointracker_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewTestDbConnection() *gorm.DB {
	dsn := "host=172.21.0.2 user=ecointracker password=jJ99268940*/ecointracker dbname=ecointracker_test_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
