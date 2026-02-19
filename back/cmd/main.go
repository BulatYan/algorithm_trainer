package main

import (
	"log"
	"os"
	"path/filepath"
	"ped_poject/config"
	"ped_poject/database"
	"ped_poject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Testing config and database...")

	// Загружаем конфиг и подключаемся к БД
	dbConfig := config.LoadDB()
	db := database.Connect(dbConfig)
	database.ApplyMigrations(db, "back/migrations")
	defer db.Close()

	userRepo := database.NewUserRepository(db)
	r := routes.SetupRouter(userRepo)

	// Получаем абсолютный путь до файла
	absPath, err := filepath.Abs("../front/register.html")
	if err != nil {
		log.Printf("Error getting absolute path: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		log.Println("Request received for /")

		// Проверяем существование файла
		_, err := os.Stat(absPath)
		if os.IsNotExist(err) {
			log.Printf("File not found: %v", err)
			c.JSON(500, gin.H{"error": "register.html not found"})
			return
		} else if err != nil {
			log.Printf("Error checking file: %v", err)
			c.JSON(500, gin.H{"error": "Failed to check register.html"})
			return
		}
		log.Printf("Absolute path to register.html: %v", absPath)

		// Отдаем файл
		c.File(absPath)
	})

	log.Println("✅ Everything works!")
	r.Run("0.0.0.0:8080")
}
