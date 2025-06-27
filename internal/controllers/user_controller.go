// internal/controllers/user_controller.go
package controllers

import (
	"net/http"
	"online-school/internal/models"
	"online-school/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Регистрация нового пользователя в системе
// @Tags public
// @Accept json
// @Produce json
// @Param input body models.User true "Данные пользователя"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := c.userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// Login godoc
// @Summary Авторизация пользователя
// @Description Вход в систему и получение JWT токена
// @Tags public
// @Accept json
// @Produce json
// @Param input body models.LoginRequest true "Данные для входа"
// @Success 200 {object} models.LoginResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var loginReq models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userService.Authenticate(loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
func (uc *UserController) GetProfile(c *gin.Context) {
	// TODO: Implement logic to get the current user's profile
	// Example placeholder:
	c.JSON(200, gin.H{"message": "GetProfile not implemented"})
}
