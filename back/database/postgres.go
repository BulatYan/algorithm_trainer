package database

import (
	"database/sql"
	"fmt"
	"log"
	"ped_poject/config"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.ConfigDB) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("postregsql connection error: %v", err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("postgresql ping error: %v", err)
	}
	log.Println("postgresql succesfully connected")
	return db
}
