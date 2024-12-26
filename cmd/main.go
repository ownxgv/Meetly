package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"meetly/config"
	_ "meetly/docs"
	"meetly/internal/users"
)

// @title Meetly API
// @version 1.0
// @description This is the API documentation for the Meetly service.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5432
// @BasePath /
func main() {
	dsn := "host=localhost user=postgres password=eternal dbname=meetly port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := config.InitDB(dsn)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	err = db.AutoMigrate(&users.User{})
	if err != nil {
		log.Fatalf("error auto migrating users: %v", err)
	}

	log.Println("Database connected and migrated successfully")

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	handler := users.NewHandler(service)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/users", func(c *gin.Context) {
		handler.GetAllUsers(&users.GinContextAdapter{C: c})
	})
	router.POST("/users", func(c *gin.Context) {
		handler.CreateUser(&users.GinContextAdapter{C: c})
	})
	router.GET("/users/:id", func(c *gin.Context) {
		handler.GetUserByID(&users.GinContextAdapter{C: c})
	})
	router.PUT("/users/:id", func(c *gin.Context) {
		handler.UpdateUser(&users.GinContextAdapter{C: c})
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		handler.DeleteUser(&users.GinContextAdapter{C: c})
	})

	log.Println("Starting server on :8080")
	router.Run(":8080")
}
