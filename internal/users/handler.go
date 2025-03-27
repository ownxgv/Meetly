package users

import (
	"meetly/internal/context"
	"net/http"
	"strconv"
)

type handler struct {
	service UserService
}

// NewHandler создаёт новый экземпляр UserHandler
func NewHandler(service UserService) UserHandler {
	return &handler{service: service}
}

// GetAllUsers - получение всех пользователей
func (h *handler) GetAllUsers(c context.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser - создание нового пользователя
func (h *handler) CreateUser(c context.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}
	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUserByID - получение пользователя по ID
func (h *handler) GetUserByID(c context.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
		return
	}
	if user == nil || user.UserID == 0 {
		c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser - обновление данных пользователя
func (h *handler) UpdateUser(c context.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}

	user.UserID = int64(id) // Привязка ID из URL

	if err := h.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser - удаление пользователя
func (h *handler) DeleteUser(c context.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
