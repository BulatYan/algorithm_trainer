package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	// JWTSecret  string
	// JWTTimeout int
}

type ConfigJWT struct {
	JWTSecret  string
	JWTTimeout int
}

func LoadDB() *ConfigDB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".env file load error: %v", err)
	}
	return &ConfigDB{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		// JWTSecret:  os.Getenv("JWT_SECRET"),
		// JWTTimeout: jwtTimeout,
	}
}

func LoadJWT() *ConfigJWT {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".env file load error: %v", err)
	}
	jwtTimeout, err := strconv.Atoi(os.Getenv("JWT_TIMEOUT"))
	if err != nil || jwtTimeout <= 0 {
		log.Println("something wrong with JWTTimeout")
		jwtTimeout = 60
	}
	return &ConfigJWT{
		JWTSecret:  os.Getenv("JWT_SECRET"),
		JWTTimeout: jwtTimeout,
	}
}
