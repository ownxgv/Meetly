package main

import (
	"log"
	"meetly/internal/app"
)

func recoverFromPanic() {
	if err := recover(); err != nil {
		log.Fatalf("❌ Critical error: %v", err)
	}
}

func main() {
	defer recoverFromPanic()

	port := "5000"
	engine := app.InitializeApp()

	log.Printf("🚀 Server is running on port %s", port)
	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
