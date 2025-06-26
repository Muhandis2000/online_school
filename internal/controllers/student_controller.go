package controllers

import (
	"net/http"

	"online-school/internal/models"
	"online-school/internal/services"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentService *services.StudentService
}

func NewStudentController(studentService *services.StudentService) *StudentController {
	return &StudentController{studentService: studentService}
}

func (c *StudentController) GetSchedule(ctx *gin.Context) {
	studentID, _ := ctx.Get("userID")
	schedules, err := c.studentService.GetSchedule(ctx.Request.Context(), studentID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, schedules)
}

func (c *StudentController) GetLessons(ctx *gin.Context) {
	studentID, _ := ctx.Get("userID")
	lessons, err := c.studentService.GetLessons(ctx.Request.Context(), studentID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, lessons)
}

func (c *StudentController) SubmitHomework(ctx *gin.Context) {
	studentID, _ := ctx.Get("userID")

	var homework models.Homework
	if err := ctx.ShouldBindJSON(&homework); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	homework.StudentID = studentID.(int)
	if err := c.studentService.SubmitHomework(ctx.Request.Context(), &homework); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, homework)
}

func (c *StudentController) GetGrades(ctx *gin.Context) {
	studentID, _ := ctx.Get("userID")
	grades, err := c.studentService.GetGrades(ctx.Request.Context(), studentID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, grades)
}

func (c *StudentController) GetAttendance(ctx *gin.Context) {
	studentID, _ := ctx.Get("userID")
	attendance, err := c.studentService.GetAttendance(ctx.Request.Context(), studentID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, attendance)
}
