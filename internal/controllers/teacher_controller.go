package controllers

import (
	"net/http"

	"online-school/internal/models"
	"online-school/internal/services"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	teacherService *services.TeacherService
}

func NewTeacherController(teacherService *services.TeacherService) *TeacherController {
	return &TeacherController{teacherService: teacherService}
}

func (c *TeacherController) CreateLesson(ctx *gin.Context) {
	teacherID, _ := ctx.Get("userID")

	var lesson models.Lesson
	if err := ctx.ShouldBindJSON(&lesson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson.TeacherID = teacherID.(int)
	if err := c.teacherService.CreateLesson(ctx.Request.Context(), &lesson); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, lesson)
}

func (c *TeacherController) RecordAttendance(ctx *gin.Context) {
	var attendance models.Attendance
	if err := ctx.ShouldBindJSON(&attendance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.teacherService.RecordAttendance(ctx.Request.Context(), &attendance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, attendance)
}

func (c *TeacherController) RecordGrade(ctx *gin.Context) {
	var grade models.Grade
	if err := ctx.ShouldBindJSON(&grade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.teacherService.RecordGrade(ctx.Request.Context(), &grade); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, grade)
}

func (c *TeacherController) GetLessons(ctx *gin.Context) {
	teacherID, _ := ctx.Get("userID")
	lessons, err := c.teacherService.GetLessonsByTeacher(ctx.Request.Context(), teacherID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, lessons)
}

func (c *TeacherController) GetSchedule(ctx *gin.Context) {
	teacherID, _ := ctx.Get("userID")
	schedules, err := c.teacherService.GetScheduleByTeacher(ctx.Request.Context(), teacherID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, schedules)
}
