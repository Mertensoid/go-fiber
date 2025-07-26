package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load("/Users/admin/Documents/Обучение/GOLANG/PurpleSchool/go-fiber/.env"); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(".env loaded")
}

type DatabaseConfig struct {
	url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}
}
