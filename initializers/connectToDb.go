package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dbc := os.Getenv("DB_URL")
	if dbc == "" {
		log.Fatal("DB_URL not set")
	}

	DB, err = gorm.Open(postgres.Open(dbc), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}
