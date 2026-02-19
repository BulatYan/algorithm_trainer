package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func ApplyMigrations(db *sql.DB, migrationsPath string) {
	// Получаем список всех файлов с расширением .sql
	files, err := filepath.Glob(filepath.Join(migrationsPath, "*.sql"))
	if err != nil {
		log.Fatalf("Error finding migration files: %v", err)
	}

	// Сортируем файлы по имени, чтобы применялись в правильном порядке
	sort.Strings(files)

	// Применяем каждую миграцию
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Error reading migration file %s: %v", file, err)
			continue // продолжаем с остальными миграциями
		}

		// Выполняем SQL запрос из файла
		_, err = db.Exec(string(content))
		if err != nil {
			log.Printf("Error applying migration %s: %v", file, err)
			continue // продолжаем с остальными миграциями
		}

		log.Printf("Migration %s successfully applied", file)
	}

	log.Println("All migrations processed")
}
