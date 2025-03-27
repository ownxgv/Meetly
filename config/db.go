package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB инициализирует подключение к БД
func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback для локальной отладки без docker
		dsn = fmt.Sprintf(
			"host=localhost user=postgres password=eternal dbname=meetly port=5432 sslmode=disable",
		)
		log.Println("⚠️  DATABASE_URL не задан, используется fallback DSN")
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к БД: %v", err)
	}

	log.Println("✅ Успешное подключение к базе данных PostgreSQL")
}

// GetDB возвращает экземпляр подключения
func GetDB() *gorm.DB {
	return db
}
