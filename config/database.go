package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbc := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbc), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	DB = db
	fmt.Println("Connected to database")
}
