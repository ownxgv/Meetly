package participants

import (
	"net/http"
	"strconv"
)

type handler struct {
	service ParticipantService
}

func NewParticipantHandler(service ParticipantService) ParticipantHandler {
	return &handler{service: service}
}

func (h *handler) GetAllParticipants(c HTTPContext) {
	participants, err := h.service.GetAllParticipants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch participants"})
		return
	}
	c.JSON(http.StatusOK, participants)
}

func (h *handler) GetParticipantsByMeetingID(c HTTPContext) {
	meetingID, err := strconv.Atoi(c.Param("meeting_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid meeting ID"})
		return
	}
	participants, err := h.service.GetParticipantsByMeetingID(uint(meetingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch participants"})
		return
	}
	c.JSON(http.StatusOK, participants)
}

func (h *handler) AddParticipant(c HTTPContext) {
	var participant Participant
	if err := c.BindJSON(&participant); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}
	if err := h.service.AddParticipant(&participant); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add participant"})
		return
	}
	c.JSON(http.StatusCreated, participant)
}

func (h *handler) UpdateParticipantStatus(c HTTPContext) {
	meetingID, err := strconv.Atoi(c.Param("meeting_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid meeting ID"})
		return
	}
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	var input struct {
		Status string `json:"status"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}
	if err := h.service.UpdateParticipantStatus(uint(meetingID), uint(userID), input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update participant status"})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Participant status updated"})
}

func (h *handler) RemoveParticipant(c HTTPContext) {
	meetingID, err := strconv.Atoi(c.Param("meeting_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid meeting ID"})
		return
	}
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	if err := h.service.RemoveParticipant(uint(meetingID), uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to remove participant"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
