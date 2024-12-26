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

type Response struct {
	Message string `json:"message"`
}

// NewHandler создаёт новый хендлер
func NewHandler(service Service) Handler {
	return &handlerForUsers{service: service}
}

// @Summary Get all users
// @Description Retrieve all users
// @Tags Users
// @Produce json
// @Success 200 {array} users.User
// @Failure 500 {object} Response
// @Router /users [get]
func (h *handlerForUsers) GetAllUsers(c Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Create a new user
// @Description Add a new user to the system
// @Tags Users
// @Accept json
// @Produce json
// @Param user body users.User true "User data"
// @Success 201 {object} users.User
// @Failure 400 {object} Response
// @Router /users [post]
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

// @Summary Get user by ID
// @Description Retrieve a single user by ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} users.User
// @Failure 404 {object} Response
// @Router /users/{id} [get]
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

// @Summary Update user
// @Description Update user details by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body users.User true "User data"
// @Success 200 {object} users.User
// @Failure 400 {object} Response
// @Router /users/{id} [put]
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

// @Summary Delete user
// @Description Remove a user by ID
// @Tags Users
// @Param id path int true "User ID"
// @Success 204 {object} nil
// @Failure 404 {object} Response
// @Router /users/{id} [delete]
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
