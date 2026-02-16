package main

import (
	"log"
	"ped_poject/config"
	"ped_poject/database"
	"ped_poject/routes"
)

func main() {
	log.Println("Testing config and database...")

	// Просто загружаем конфиг и подключаемся к БД
	dbConfig := config.LoadDB()
	db := database.Connect(dbConfig)
	database.ApplyMigrations(db, "internal/migrations")
	defer db.Close()

	userRepo := database.NewUserRepository(db)
	r := routes.SetupRouter(userRepo)

	log.Println("✅ Everything works!")
	r.Run("0.0.0.0:8080")
}
