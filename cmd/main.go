package main

import (
	"log"
	"meetly/config"
	"meetly/internal/context"
	"meetly/internal/meetings"
	"meetly/internal/participants"
	"meetly/internal/users"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// recoverFromPanic предотвращает краш сервера
func recoverFromPanic() {
	if err := recover(); err != nil {
		log.Fatalf("❌ Critical error: %v", err)
	}
}

func main() {
	defer recoverFromPanic()

	// Бэкенд работает на 5000 порту (совместимость с фронтом)
	port := "5000"

	// Подключаем базу данных
	config.InitDB()
	if config.DB == nil {
		log.Fatalf("❌ Failed to connect to the database. Server stopped.")
	}
	log.Println("✅ Successfully connected to the database!")

	// Применяем миграции (таблицы создадутся автоматически)
	log.Println("🛠️ Running AutoMigrate...")
	err := config.DB.AutoMigrate(&users.User{}, &meetings.Meeting{}, &participants.Participant{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migrations applied successfully!")

	// Настраиваем сервер Gin
	router := gin.Default()

	// Включаем CORS (чтобы фронтенд мог делать запросы)
	router.Use(cors.Default())

	// 📌 **Маршруты пользователей**
	userRepo := users.NewRepository(config.DB)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	router.GET("/users", func(c *gin.Context) {
		userHandler.GetAllUsers(&context.GinContextAdapter{C: c})
	})
	router.POST("/users", func(c *gin.Context) {
		userHandler.CreateUser(&context.GinContextAdapter{C: c})
	})
	router.GET("/users/:id", func(c *gin.Context) {
		userHandler.GetUserByID(&context.GinContextAdapter{C: c})
	})
	router.PUT("/users/:id", func(c *gin.Context) {
		userHandler.UpdateUser(&context.GinContextAdapter{C: c})
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		userHandler.DeleteUser(&context.GinContextAdapter{C: c})
	})

	// 📌 **Маршруты встреч (meetings)**
	meetingRepo := meetings.NewMeetingRepository(config.DB)
	meetingService := meetings.NewMeetingService(meetingRepo)
	meetingHandler := meetings.NewHandler(meetingService)

	router.GET("/meetings", func(c *gin.Context) {
		meetingHandler.GetAllMeetings(&context.GinContextAdapter{C: c})
	})
	router.POST("/meetings", func(c *gin.Context) {
		meetingHandler.CreateMeeting(&context.GinContextAdapter{C: c})
	})
	router.GET("/meetings/:id", func(c *gin.Context) {
		meetingHandler.GetMeetingByID(&context.GinContextAdapter{C: c})
	})
	router.PUT("/meetings/:id", func(c *gin.Context) {
		meetingHandler.UpdateMeeting(&context.GinContextAdapter{C: c})
	})
	router.DELETE("/meetings/:id", func(c *gin.Context) {
		meetingHandler.DeleteMeeting(&context.GinContextAdapter{C: c})
	})

	// 📌 **Маршруты участников (participants)**
	participantRepo := participants.NewParticipantRepository(config.DB)
	participantService := participants.NewParticipantService(participantRepo)
	participantHandler := participants.NewParticipantHandler(participantService)

	router.GET("/participants", func(c *gin.Context) {
		participantHandler.GetAllParticipants(&context.GinContextAdapter{C: c})
	})
	router.POST("/participants", func(c *gin.Context) {
		participantHandler.AddParticipant(&context.GinContextAdapter{C: c})
	})
	router.GET("/participants/:id", func(c *gin.Context) {
		participantHandler.GetParticipantsByMeetingID(&context.GinContextAdapter{C: c})
	})
	router.PUT("/participants/:id", func(c *gin.Context) {
		participantHandler.UpdateParticipantStatus(&context.GinContextAdapter{C: c})
	})
	router.DELETE("/participants/:id", func(c *gin.Context) {
		participantHandler.RemoveParticipant(&context.GinContextAdapter{C: c})
	})

	// 📌 **Маршрут `/events/:id`, который сразу возвращает всю информацию**
	router.GET("/events/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Получаем событие и сразу подтягиваем `Creator`
		var event meetings.Meeting
		if err := config.DB.Preload("Creator").First(&event, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}

		// Получаем участников
		var participants []participants.Participant
		config.DB.Where("meeting_id = ?", id).Find(&participants)

		// Преобразуем участников в JSON-формат
		var participantList []gin.H
		for _, p := range participants {
			var user users.User
			config.DB.First(&user, p.UserID)
			participantList = append(participantList, gin.H{
				"id":   user.UserID,
				"name": user.Username,
			})
		}

		// ✅ Теперь можно использовать `event.Creator`
		c.JSON(http.StatusOK, gin.H{
			"id":          event.ID,
			"title":       event.Title,
			"date":        event.MeetingTime,
			"description": event.Description,
			"address":     event.Location,
			"creator": gin.H{
				"id":   event.CreatorID, // Теперь берем из `event.Creator`
				"name": event.Creator,
			},
			"participants": participantList,
		})
	})

	// Проверочный маршрут `/health`
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Запуск сервера на 5000 порту
	log.Printf("🚀 Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Error starting the server: %v", err)
	}
}
