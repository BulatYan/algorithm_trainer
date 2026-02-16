package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

func ApplyMigrations(db *sql.DB, migrationsPath string) {
	files, err := filepath.Glob(filepath.Join(migrationsPath, "*.sql"))
	if err != nil {
		log.Fatalf("migrations file error: %v", err)
	}
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("reading %v error: %v", file, err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatalf("applying %v error: %v", file, err)
		}
		log.Printf("migration %v succesfully applied\n", file)
	}
	log.Println("all migrations succesfully applied")
}
