package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Muhandis2000/online-school/internal/config"
	"github.com/Muhandis2000/online-school/internal/controllers"
	"github.com/Muhandis2000/online-school/internal/db"
	"github.com/Muhandis2000/online-school/internal/middleware"
	"github.com/Muhandis2000/online-school/internal/repositories"
	"github.com/Muhandis2000/online-school/internal/services"
	"github.com/gin-gonic/gin"
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

	// 4. Инициализация слоев приложения
	userRepo := repositories.NewUserRepository(dbConn.DB)
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret)
	authController := controllers.NewAuthController(authService)

	// 5. Настройка роутера
	router := gin.Default()

	// Публичные маршруты
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	// Защищенные маршруты
	authGroup := router.Group("/")
	authGroup.Use(middleware.JWTAuthMiddleware(cfg.JWT.Secret))
	{
		// Здесь будут защищенные эндпоинты
		authGroup.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Это защищенный ресурс!"})
		})
	}

	// 6. Запуск сервера
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

	// 7. Обработка сигналов для graceful shutdown
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
