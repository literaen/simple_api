package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "host=192.168.50.154 user=postgres password=yourpassword dbname=postgres port=5432"
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database", err)
	}
}
