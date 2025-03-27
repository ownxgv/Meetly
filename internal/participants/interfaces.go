package participants

import (
	"github.com/gin-gonic/gin"
	"meetly/internal/context"
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
	UpdateParticipantStatus(meetingID uint, userID uint, status string) error
	RemoveParticipant(meetingID uint, userID uint) error
}

type ParticipantHandler interface {
	GetAllParticipants(c context.Context)
	GetParticipantsByMeetingID(c context.Context)
	AddParticipant(c context.Context)
	UpdateParticipantStatus(c context.Context)
	RemoveParticipant(c context.Context)
}

//type Context interface {
//	JSON(code int, obj interface{}) error
//	BindJSON(obj interface{}) error
//	Param(key string) string
//}

// GinContextAdapter — адаптер для использования Gin в интерфейсе HTTPContext
type GinContextAdapter struct {
	C *gin.Context
}
