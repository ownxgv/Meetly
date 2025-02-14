package main

import (
	"fmt"
	"log"
	"meetly/config"
	"meetly/internal/context"
	"meetly/internal/meetings"
	"meetly/internal/participants"
	"meetly/internal/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

// recoverFromPanic catches any panics and prevents the server from crashing
func recoverFromPanic() {
	if err := recover(); err != nil {
		log.Fatalf("❌ Critical error: %v", err)
	}
}

func main() {
	// Catch unexpected panics
	defer recoverFromPanic()

	// Check if the port is already in use
	port := "8080"
	if isPortInUse(port) {
		log.Fatalf("❌ Port %s is already in use. Try a different port.", port)
	}

	// Initialize the database
	log.Println("🔄 Connecting to the database...")
	config.InitDB()
	if config.DB == nil {
		log.Fatalf("❌ Failed to connect to the database. Server stopped.")
	}
	log.Println("✅ Successfully connected to the database!")

	// Run database migrations
	log.Println("🛠️ Running AutoMigrate...")
	err := config.DB.AutoMigrate(&users.User{}, &meetings.Meeting{}, &participants.Participant{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migrations applied successfully!")

	// 💾 Создаём объект database для передачи в сервисы
	database := config.DB

	// Initialize dependencies
	userRepo := users.NewRepository(database)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	meetingRepo := meetings.NewMeetingRepository(database)
	meetingService := meetings.NewMeetingService(meetingRepo)
	meetingHandler := meetings.NewHandler(meetingService)

	participantRepo := participants.NewParticipantRepository(database)
	participantService := participants.NewParticipantService(participantRepo)
	participantHandler := participants.NewParticipantHandler(participantService)

	// Set up routes
	router := gin.Default()

	// 📌 User routes
	router.GET("/users", func(c *gin.Context) {
		userHandler.GetAllUsers(&context.GinContextAdapter{C: c})
	})
	router.POST("/users", func(c *gin.Context) {
		userHandler.CreateUser(&context.GinContextAdapter{C: c})
	})

	// 📌 Meeting routes
	router.GET("/meetings", func(c *gin.Context) {
		meetingHandler.GetAllMeetings(&context.GinContextAdapter{C: c})
	})
	router.POST("/meetings", func(c *gin.Context) {
		meetingHandler.CreateMeeting(&context.GinContextAdapter{C: c})
	})

	// 📌 Participant routes
	router.GET("/participants", func(c *gin.Context) {
		participantHandler.GetAllParticipants(&context.GinContextAdapter{C: c})
	})
	router.POST("/participants", func(c *gin.Context) {
		participantHandler.AddParticipant(&context.GinContextAdapter{C: c})
	})

	// 📌 Start the server with error handling
	log.Printf("🚀 Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Error starting the server: %v", err)
	}
}

// isPortInUse checks if the specified port is already occupied
func isPortInUse(port string) bool {
	conn, err := http.Get(fmt.Sprintf("http://localhost:%s", port))
	if err == nil {
		conn.Body.Close()
		return true
	}
	return false
}
