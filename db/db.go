package db

import (
	"log"

	"github.com/OliveiraJ/ecoin-tracker-v2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	err := db.AutoMigrate(&models.Read{})
	if err != nil {
		panic(err)
	}
}
func NewDbConnection() *gorm.DB {
	dsn := "host=172.20.0.3 user=ecointracker password=jJ99268940*/ecointracker dbname=ecointracker_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
