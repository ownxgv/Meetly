package main

import (
	"log"
	"meetly/internal/meetings"
	"meetly/internal/participants"
	"meetly/internal/users"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Подключение к базе данных
	dsn := "host=localhost user=postgres password=yourpassword dbname=meetly port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Автоматические миграции
	db.AutoMigrate(&users.User{}, &meetings.Meeting{}, &participants.Participant{})

	// Инициализация зависимостей для пользователей
	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	// Инициализация зависимостей для встреч
	meetingRepo := meetings.NewMeetingRepository(db)
	meetingService := meetings.NewMeetingService(meetingRepo)
	meetingHandler := meetings.NewHandler(meetingService)

	// Инициализация зависимостей для участников
	participantRepo := participants.NewParticipantRepository(db)
	participantService := participants.NewParticipantService(participantRepo)
	participantHandler := participants.NewParticipantHandler(participantService)

	// Создание роутера Gin
	router := gin.Default()

	// Маршруты для пользователей
	router.GET("/users", func(c *gin.Context) {
		userHandler.GetAllUsers(&users.GinContextAdapter{C: c})
	})
	router.POST("/users", func(c *gin.Context) {
		userHandler.CreateUser(&users.GinContextAdapter{C: c})
	})
	router.GET("/users/:id", func(c *gin.Context) {
		userHandler.GetUserByID(&users.GinContextAdapter{C: c})
	})
	router.PUT("/users/:id", func(c *gin.Context) {
		userHandler.UpdateUser(&users.GinContextAdapter{C: c})
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		userHandler.DeleteUser(&users.GinContextAdapter{C: c})
	})

	// Маршруты для встреч
	router.GET("/meetings", func(c *gin.Context) {
		meetingHandler.GetAllMeetings(&meetings.GinContextAdapter{C: c})
	})
	router.POST("/meetings", func(c *gin.Context) {
		meetingHandler.CreateMeeting(&meetings.GinContextAdapter{C: c})
	})
	router.GET("/meetings/:id", func(c *gin.Context) {
		meetingHandler.GetMeetingByID(&meetings.GinContextAdapter{C: c})
	})
	router.PUT("/meetings/:id", func(c *gin.Context) {
		meetingHandler.UpdateMeeting(&meetings.GinContextAdapter{C: c})
	})
	router.DELETE("/meetings/:id", func(c *gin.Context) {
		meetingHandler.DeleteMeeting(&meetings.GinContextAdapter{C: c})
	})

	// Маршруты для участников
	router.GET("/participants", func(c *gin.Context) {
		participantHandler.GetAllParticipants(&participants.GinContextAdapter{C: c})
	})
	router.GET("/participants/meeting/:meeting_id", func(c *gin.Context) {
		participantHandler.GetParticipantsByMeetingID(&participants.GinContextAdapter{C: c})
	})
	router.POST("/participants", func(c *gin.Context) {
		participantHandler.AddParticipant(&participants.GinContextAdapter{C: c})
	})
	router.PUT("/participants/meeting/:meeting_id/user/:user_id", func(c *gin.Context) {
		participantHandler.UpdateParticipantStatus(&participants.GinContextAdapter{C: c})
	})
	router.DELETE("/participants/meeting/:meeting_id/user/:user_id", func(c *gin.Context) {
		participantHandler.RemoveParticipant(&participants.GinContextAdapter{C: c})
	})

	// Запуск сервера
	log.Println("Server running on port 8080")
	router.Run(":8080")
}
