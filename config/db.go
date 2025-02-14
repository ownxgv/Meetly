package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDB подключает базу данных и сохраняет её в глобальной переменной
func InitDB() {
	dsn := "host=localhost user=postgres password=eternal dbname=meetly port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error connecting DB: %v", err)
	}
	log.Println("✅ Connected successfully PostgresSQL")
}
