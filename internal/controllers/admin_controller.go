// CreateTeacher godoc
// @Summary Create a new teacher
// @Description Create a new teacher (admin only)
// @Tags admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   teacher body     models.Teacher true "Teacher Data"
// @Success 201     {object} models.Teacher
// @Failure 400     {object} map[string]string
// @Failure 403     {object} map[string]string
// @Failure 500     {object} map[string]string
// @Router  /admin/teachers [post]
package controllers

import (
	"net/http"
	"strconv"

	"online-school/internal/models"
	"online-school/internal/services"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *services.AdminService
}

func NewAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

func (c *AdminController) CreateTeacher(ctx *gin.Context) {
	var teacher models.Teacher
	if err := ctx.ShouldBindJSON(&teacher); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.adminService.CreateTeacher(ctx.Request.Context(), &teacher); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, teacher)
}

func (c *AdminController) ApproveTeacher(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid teacher ID"})
		return
	}

	if err := c.adminService.ApproveTeacher(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "teacher approved"})
}

func (c *AdminController) CreateStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.adminService.CreateStudent(ctx.Request.Context(), &student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, student)
}

func (c *AdminController) CreateLesson(ctx *gin.Context) {
	var lesson models.Lesson
	if err := ctx.ShouldBindJSON(&lesson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.adminService.CreateLesson(ctx.Request.Context(), &lesson); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, lesson)
}

func (c *AdminController) CreateSchedule(ctx *gin.Context) {
	var schedule models.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.adminService.CreateSchedule(ctx.Request.Context(), &schedule); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, schedule)
}

func (c *AdminController) CreatePayment(ctx *gin.Context) {
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.adminService.CreatePayment(ctx.Request.Context(), &payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, payment)
}
