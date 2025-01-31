package participants

import "time"

type Participant struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"participant_id"`
	MeetingID uint      `gorm:"not null" json:"meeting_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Status    string    `gorm:"size:20" json:"status"`
	JoinedAt  time.Time `gorm:"autoCreateTime" json:"joined_at"`
}
