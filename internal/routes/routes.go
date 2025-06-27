package routes

import (
	"online-school/internal/controllers"
	"online-school/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Online School API
// @version 1.0
// @description API для управления онлайн-школой
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetupRouter(
	adminController *controllers.AdminController,
	teacherController *controllers.TeacherController,
	studentController *controllers.StudentController,
) *gin.Engine {
	router := gin.Default()

	// Swagger документация
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группа публичных эндпоинтов

	// Группа приватных эндпоинтов с авторизацией
	private := router.Group("/api/v1")
	private.Use(middleware.AuthMiddleware())

	// Админские эндпоинты
	admin := private.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	{
		// Создание преподавателя
		// @Summary Создать запись преподавателя
		// @Description Создание нового преподавателя (только для админов)
		// @Tags admin
		// @Security BearerAuth
		// @Accept json
		// @Produce json
		// @Param input body models.Teacher true "Данные преподавателя"
		// @Success 201 {object} models.Teacher
		// @Failure 403 {object} models.ErrorResponse
		// @Router /admin/teachers [post]
		admin.POST("/teachers", adminController.CreateTeacher)

		// Создание студента
		// @Summary Создать запись студента
		// @Description Создание нового студента (только для админов)
		// @Tags admin
		// @Security BearerAuth
		// @Accept json
		// @Produce json
		// @Param input body models.Student true "Данные студента"
		// @Success 201 {object} models.Student
		// @Failure 403 {object} models.ErrorResponse
		// @Router /admin/students [post]
		admin.POST("/students", adminController.CreateStudent)
	}

	// Эндпоинты преподавателя
	teacher := private.Group("/teacher")
	// teacher.Use(middleware.TeacherMiddleware()) // TODO: Implement TeacherMiddleware in middleware package
	{
		// Создание урока
		// @Summary Создать новый урок
		// @Description Создание урока (только для преподавателей)
		// @Tags teacher
		// @Security BearerAuth
		// @Accept json
		// @Produce json
		// @Param input body models.Lesson true "Данные урока"
		// @Success 201 {object} models.Lesson
		// @Failure 403 {object} models.ErrorResponse
		// @Router /teacher/lessons [post]
		teacher.POST("/lessons", teacherController.CreateLesson)
	}

	// Эндпоинты студента
	student := private.Group("/student")
	{
		// Получение расписания
		// @Summary Получить расписание занятий
		// @Description Получение расписания для текущего студента
		// @Tags student
		// @Security BearerAuth
		// @Produce json
		// @Success 200 {array} models.Schedule
		// @Failure 403 {object} models.ErrorResponse
		// @Router /student/schedule [get]
		student.GET("/schedule", studentController.GetSchedule)
	}

	return router
}
