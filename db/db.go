package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Event struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func InitPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Event{})
}

func CreateEvent(event *Event) (*Event, error) {
	event.ID = uuid.New().String()
	res := db.Create(&event)
	if res.Error != nil {
		return nil, res.Error
	}
	return event, nil
}

func GetEvent(id string) (*Event, error) {
	var event Event
	res := db.First(&event, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("event of id %s not found", id))
	}
	return &event, nil
}

func GetEvents() ([]*Event, error) {
	var events []*Event
	res := db.Find(&events)
	if res.Error != nil {
		return nil, errors.New("no events found")
	}
	return events, nil
}

func UpdateEvent(event *Event) (*Event, error) {
	var eventToUpdate Event
	result := db.Model(&eventToUpdate).Where("id = ?", event.ID).Updates(event)
	if result.RowsAffected == 0 {
		return &eventToUpdate, errors.New("event not updated")
	}
	return event, nil
}

func DeleteEvent(id string) error {
	var deletedEvent Event
	result := db.Where("id = ?", id).Delete(&deletedEvent)
	if result.RowsAffected == 0 {
		return errors.New("event not deleted")
	}
	return nil
}
