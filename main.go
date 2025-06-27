// @title Online School API
// @version 1.0
// @description API для системы управления онлайн-школой

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"online-school/internal/config"
	"online-school/internal/controllers"
	"online-school/internal/db"
	"online-school/internal/middleware"
	"online-school/internal/repositories"
	"online-school/internal/services"

	_ "online-school/docs" // Импортируем сгенерированную документацию

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 1. Загрузка конфигурации
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("❌ Не удалось загрузить конфигурацию: %v", err)
	}
	log.Println("✅ Конфигурация загружена")

	// 2. Подключение к базе данных
	dbConn, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к базе данных: %v", err)
	}
	defer dbConn.Close()
	log.Println("✅ Подключение к базе данных установлено")

	// 3. Применение миграций
	if err := db.RunMigrations(dbConn.DB); err != nil {
		log.Fatalf("❌ Ошибка миграций: %v", err)
	}
	log.Println("✅ Миграции базы данных применены")

	// 4. Инициализация репозиториев
	userRepo := repositories.NewUserRepository(dbConn.DB)
	teacherRepo := repositories.NewTeacherRepository(dbConn.DB)
	studentRepo := repositories.NewStudentRepository(dbConn.DB)
	lessonRepo := repositories.NewLessonRepository(dbConn.DB)
	scheduleRepo := repositories.NewScheduleRepository(dbConn.DB)
	homeworkRepo := repositories.NewHomeworkRepository(dbConn.DB)
	gradeRepo := repositories.NewGradeRepository(dbConn.DB)
	attendanceRepo := repositories.NewAttendanceRepository(dbConn.DB)
	paymentRepo := repositories.NewPaymentRepository(dbConn.DB)

	// 5. Инициализация сервисов
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret)
	adminService := services.NewAdminService(
		userRepo,
		teacherRepo,
		studentRepo,
		lessonRepo,
		scheduleRepo,
		paymentRepo,
	)
	teacherService := services.NewTeacherService(
		teacherRepo,
		lessonRepo,
		scheduleRepo,
		attendanceRepo,
		gradeRepo,
	)
	studentService := services.NewStudentService(
		studentRepo,
		scheduleRepo,
		lessonRepo,
		homeworkRepo,
		gradeRepo,
		attendanceRepo,
	)

	// 6. Инициализация контроллеров
	authController := controllers.NewAuthController(authService)
	adminController := controllers.NewAdminController(adminService)
	teacherController := controllers.NewTeacherController(teacherService)
	studentController := controllers.NewStudentController(studentService)

	// 7. Настройка Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Добавляем роут для Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Настройка CORS
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	// Настройка логирования
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 8. Настройка маршрутов
	// Публичные маршруты
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	// Защищенные маршруты
	authGroup := router.Group("/")
	authGroup.Use(middleware.JWTAuthMiddleware(cfg.JWT.Secret))
	{
		// Административные функции
		adminGroup := authGroup.Group("/admin")
		adminGroup.Use(middleware.RoleMiddleware("admin"))
		{
			adminGroup.POST("/teachers", adminController.CreateTeacher)
			adminGroup.PUT("/teachers/:id/approve", adminController.ApproveTeacher)
			adminGroup.POST("/students", adminController.CreateStudent)
			adminGroup.POST("/lessons", adminController.CreateLesson)
			adminGroup.POST("/schedules", adminController.CreateSchedule)
			adminGroup.POST("/payments", adminController.CreatePayment)
		}

		// Функции преподавателя
		teacherGroup := authGroup.Group("/teacher")
		teacherGroup.Use(middleware.RoleMiddleware("teacher"))
		{
			teacherGroup.POST("/lessons", teacherController.CreateLesson)
			teacherGroup.POST("/attendance", teacherController.RecordAttendance)
			teacherGroup.POST("/grades", teacherController.RecordGrade)
			teacherGroup.GET("/lessons", teacherController.GetLessons)
			teacherGroup.GET("/schedule", teacherController.GetSchedule)
		}

		// Функции студента
		studentGroup := authGroup.Group("/student")
		studentGroup.Use(middleware.RoleMiddleware("student"))
		{
			studentGroup.GET("/schedule", studentController.GetSchedule)
			studentGroup.GET("/lessons", studentController.GetLessons)
			studentGroup.POST("/homework", studentController.SubmitHomework)
			studentGroup.GET("/grades", studentController.GetGrades)
			studentGroup.GET("/attendance", studentController.GetAttendance)
		}

		// Общие защищенные функции
		authGroup.GET("/protected", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			userRole, _ := c.Get("userRole")
			c.JSON(200, gin.H{
				"message": "Это защищенный ресурс!",
				"user_id": userID,
				"role":    userRole,
			})
		})
	}

	// 9. Запуск сервера
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	go func() {
		log.Printf("🚀 Сервер запущен на порту %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Ошибка сервера: %s\n", err)
		}
	}()

	// 10. Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("⏳ Выключение сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("❌ Принудительное выключение сервера:", err)
	}

	log.Println("✅ Сервер успешно остановлен")
}
