package router

import (
	ctx "meetly/internal/context"
	"meetly/internal/meetings"
	"meetly/internal/participants"
	"meetly/internal/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	userHandler := users.NewHandler(users.NewService(users.NewRepository(db)))
	meetingHandler := meetings.NewHandler(meetings.NewMeetingService(meetings.NewMeetingRepository(db)))
	participantHandler := participants.NewParticipantHandler(participants.NewParticipantService(participants.NewParticipantRepository(db)))

	// Users
	r.GET("/users", wrap(userHandler.GetAllUsers))
	r.POST("/users", wrap(userHandler.CreateUser))
	r.GET("/users/:id", wrap(userHandler.GetUserByID))
	r.PUT("/users/:id", wrap(userHandler.UpdateUser))
	r.DELETE("/users/:id", wrap(userHandler.DeleteUser))

	// Meetings
	r.GET("/meetings", wrap(meetingHandler.GetAllMeetings))
	r.POST("/meetings", wrap(meetingHandler.CreateMeeting))
	r.GET("/meetings/:id", wrap(meetingHandler.GetMeetingByID))
	r.PUT("/meetings/:id", wrap(meetingHandler.UpdateMeeting))
	r.DELETE("/meetings/:id", wrap(meetingHandler.DeleteMeeting))

	// Participants
	r.GET("/participants", wrap(participantHandler.GetAllParticipants))
	r.POST("/participants", wrap(participantHandler.AddParticipant))
	r.GET("/participants/:id", wrap(participantHandler.GetParticipantsByMeetingID))
	r.PUT("/participants/:id", wrap(participantHandler.UpdateParticipantStatus))
	r.DELETE("/participants/:id", wrap(participantHandler.RemoveParticipant))

	// `/events/:id` endpoint
	r.GET("/events/:id", func(c *gin.Context) {
		id := c.Param("id")
		var event meetings.Meeting
		if err := db.Preload("Creator").First(&event, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}

		var participantsList []participants.Participant
		db.Where("meeting_id = ?", id).Find(&participantsList)

		var jsonList []gin.H
		for _, p := range participantsList {
			var u users.User
			db.First(&u, p.UserID)
			jsonList = append(jsonList, gin.H{"id": u.UserID, "name": u.Username})
		}

		c.JSON(http.StatusOK, gin.H{
			"id":          event.ID,
			"title":       event.Title,
			"date":        event.MeetingTime,
			"description": event.Description,
			"address":     event.Location,
			"creator": gin.H{
				"id":   event.CreatorID,
				"name": event.Creator,
			},
			"participants": jsonList,
		})
	})

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

// Обёртка адаптера
func wrap(f func(c ctx.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(&ctx.GinContextAdapter{C: c})
	}
}
