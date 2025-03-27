package app

import (
	"log"
	"meetly/config"
	"meetly/internal/meetings"
	"meetly/internal/participants"
	"meetly/internal/router"
	"meetly/internal/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeApp() *gin.Engine {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using defaults")
	}

	// Инициализируем БД
	config.InitDB()
	db := config.GetDB()

	if db == nil {
		log.Fatal("❌ Failed to connect to DB")
	}
	log.Println("✅ Successfully connected to DB!")

	// Применяем миграции
	if err := db.AutoMigrate(
		&users.User{},
		&meetings.Meeting{},
		&participants.Participant{},
	); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migrations applied")

	// Настраиваем сервер Gin
	engine := gin.Default()
	engine.Use(cors.Default())

	// Регистрируем маршруты
	router.RegisterRoutes(engine, db)

	return engine
}
