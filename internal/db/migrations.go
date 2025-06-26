package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

var migrationsFS embed.FS

func RunMigrations(db *sqlx.DB) error {
	files, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	// Создаем таблицу для отслеживания миграций
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		var exists bool
		err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM migrations WHERE name = $1)", file.Name())
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if exists {
			log.Printf("Миграция %s уже применена, пропускаем", file.Name())
			continue
		}

		content, err := fs.ReadFile(migrationsFS, "migrations/"+file.Name())
		if err != nil {
			return err
		}

		// Выполняем миграцию
		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("ошибка применения миграции %s: %w", file.Name(), err)
		}

		// Записываем в историю миграций
		_, err = db.Exec("INSERT INTO migrations (name) VALUES ($1)", file.Name())
		if err != nil {
			return err
		}

		log.Printf("Применена миграция: %s", file.Name())
	}

	return nil
}
