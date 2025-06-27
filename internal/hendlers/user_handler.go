package hendler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService UserService
}

// UserService should be an interface or struct that provides user-related methods.
// Define UserService interface or import it from the appropriate package if not already defined.

// UserService defines the methods that the user service must implement.
type UserService interface {
	GetAllUsers() ([]interface{}, error)     // Replace interface{} with your actual User type
	GetUserByID(id int) (interface{}, error) // Replace interface{} with your actual User type
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
