package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GinContextAdapter struct {
	C *gin.Context
}

func (g *GinContextAdapter) JSON(code int, obj interface{}) error {
	g.C.JSON(code, obj)
	return nil
}

func (g *GinContextAdapter) BindJSON(obj interface{}) error {
	return g.C.ShouldBindJSON(obj)
}

type handlerForUsers struct {
	service Service
}

// NewHandler создаёт новый хендлер
func NewHandler(service Service) Handler {
	return &handlerForUsers{service: service}
}

func (h *handlerForUsers) GetAllUsers(c Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *handlerForUsers) CreateUser(c Context) {
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

func (h *handlerForUsers) GetUserByID(c Context) {
	id, err := strconv.Atoi(c.(*GinContextAdapter).C.Param("id"))
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

func (h *handlerForUsers) UpdateUser(c Context) {
	id, err := strconv.Atoi(c.(*GinContextAdapter).C.Param("id"))
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

func (h *handlerForUsers) DeleteUser(c Context) {
	id, err := strconv.Atoi(c.(*GinContextAdapter).C.Param("id"))
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
