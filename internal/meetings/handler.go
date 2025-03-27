package meetings

import (
	"meetly/internal/context"
	"net/http"
	"strconv"
)

type handler struct {
	service MeetingService
}

func NewHandler(service MeetingService) MeetingHandler {
	return &handler{service: service}
}

func (h *handler) GetAllMeetings(c context.Context) {
	meetings, err := h.service.GetAllMeetings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch meetings"})
		return
	}
	c.JSON(http.StatusOK, meetings)
}

func (h *handler) CreateMeeting(c context.Context) {
	var meeting Meeting
	if err := c.BindJSON(&meeting); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}
	if err := h.service.CreateMeeting(&meeting); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create meeting"})
		return
	}
	c.JSON(http.StatusCreated, meeting)
}

func (h *handler) GetMeetingByID(c context.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid meeting ID"})
		return
	}

	meeting, err := h.service.GetMeetingByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch meeting"})
		return
	}

	// Проверка на отсутствие
	if meeting == nil || meeting.ID == 0 {
		c.JSON(http.StatusNotFound, map[string]string{"error": "Meeting not found"})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// UpdateMeeting обновляет данные встречи
func (h *handler) UpdateMeeting(c context.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid meeting ID"})
		return
	}

	var meeting Meeting
	if err := c.BindJSON(&meeting); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}

	meeting.ID = id // Привязка ID из URL

	if err := h.service.UpdateMeeting(&meeting); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update meeting"})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// DeleteMeeting удаляет встречу по ID
func (h *handler) DeleteMeeting(c context.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid meeting ID"})
		return
	}

	if err := h.service.DeleteMeeting(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete meeting"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
