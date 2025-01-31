package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to PostgreSQL")
	return db, nil
}

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=yourpassword dbname=meetly port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
