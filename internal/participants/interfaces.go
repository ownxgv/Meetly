package participants

import (
	"github.com/gin-gonic/gin"
)

type ParticipantRepository interface {
	GetAllParticipants() ([]Participant, error)
	GetParticipantsByMeetingID(meetingID uint) ([]Participant, error)
	AddParticipant(participant *Participant) error
	UpdateParticipantStatus(meetingID, userID uint, status string) error
	RemoveParticipant(meetingID, userID uint) error
}

type ParticipantService interface {
	GetAllParticipants() ([]Participant, error)
	GetParticipantsByMeetingID(meetingID uint) ([]Participant, error)
	AddParticipant(participant *Participant) error
	UpdateParticipantStatus(meetingID, userID uint, status string) error
	RemoveParticipant(meetingID, userID uint) error
}

type ParticipantHandler interface {
	GetAllParticipants(c HTTPContext)
	GetParticipantsByMeetingID(c HTTPContext)
	AddParticipant(c HTTPContext)
	UpdateParticipantStatus(c HTTPContext)
	RemoveParticipant(c HTTPContext)
}

type HTTPContext interface {
	JSON(code int, obj interface{}) error
	BindJSON(obj interface{}) error
	Param(key string) string
}

// GinContextAdapter — адаптер для использования Gin в интерфейсе HTTPContext
type GinContextAdapter struct {
	C *gin.Context
}
