package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "db"
	port     = "5432"
	user     = "myuser"
	password = "mypassword"
	name     = "mydb"
)

func ConnectDB() error {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return fmt.Errorf("failed to connect to database")
	}

	fmt.Println("Database connection successfully")
	return nil
}
