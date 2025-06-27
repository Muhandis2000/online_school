package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS // Добавлена директива embed

func RunMigrations(db *sqlx.DB) error {
	log.Println("Запуск миграций...")

	// Читаем файлы из встроенной FS
	files, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("ошибка чтения директории миграций: %w", err)
	}

	// Сортируем файлы по имени
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	log.Printf("Найдено %d файлов миграций", len(files))

	// Создаем таблицу для отслеживания миграций
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы миграций: %w", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		log.Printf("Обработка миграции: %s", file.Name())

		// Проверяем применение миграции
		var exists bool
		err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM migrations WHERE name = $1)", file.Name())
		if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("ошибка проверки миграции: %w", err)
		}

		if exists {
			log.Printf("Миграция %s уже применена, пропускаем", file.Name())
			continue
		}

		// Читаем содержимое файла
		content, err := fs.ReadFile(migrationsFS, "migrations/"+file.Name())
		if err != nil {
			return fmt.Errorf("ошибка чтения файла миграции: %w", err)
		}

		// Выполняем миграцию
		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("ошибка выполнения миграции %s: %w", file.Name(), err)
		}

		// Регистрируем миграцию
		_, err = db.Exec("INSERT INTO migrations (name) VALUES ($1)", file.Name())
		if err != nil {
			return fmt.Errorf("ошибка регистрации миграции: %w", err)
		}

		log.Printf("✅ Применена миграция: %s", file.Name())
	}

	return nil
}
