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
func (h *handler) GetAllUsers(c context.HTTPContext) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser - создание нового пользователя
func (h *handler) CreateUser(c context.HTTPContext) {
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
func (h *handler) GetUserByID(c context.HTTPContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser - обновление данных пользователя
func (h *handler) UpdateUser(c context.HTTPContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}
	user.UserID = int64(id)
	if err := h.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser - удаление пользователя
func (h *handler) DeleteUser(c context.HTTPContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}
	if err := h.service.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
