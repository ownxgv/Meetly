package users

import "time"

type User struct {
	UserID       int64     `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username     string    `gorm:"size:50;not null" json:"username"`
	Email        string    `gorm:"size:100;unique;not null" json:"email"`
	PasswordHash string    `gorm:"size:225;not null" json:"-"`
	FullName     string    `gorm:"size:100" json:"full_name"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}
