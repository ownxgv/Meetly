package meetings

import (
	"meetly/internal/users"
	"time"
)

type Meeting struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"meeting_id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Location    string    `gorm:"size:200" json:"location"`
	MeetingTime time.Time `gorm:"not null" json:"meeting_time"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	AgeLimit    int64     `gorm:"default:0" json:"age_limit"`
	// Связь с таблицей users
	CreatorID int        `gorm:"not null" json:"creator_id"`
	Creator   users.User `gorm:"foreignKey:CreatorID" json:"creator"`
}
