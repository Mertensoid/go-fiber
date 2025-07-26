package config

import (
	"log"
	"os"
	"strconv"

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

func getString(s, defaultValue string) string {
	res := os.Getenv(s)
	if res == "" {
		return defaultValue
	} else {
		return s
	}
}

func getInt(s string, defaultValue int) int {
	res, err := strconv.Atoi(os.Getenv(s))
	if err != nil {
		return defaultValue
	}
	return res
}

func getBool(s string, defaultValue bool) bool {
	res, err := strconv.ParseBool(os.Getenv(s))
	if err != nil {
		return defaultValue
	}
	return res
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: getString("DATABASE_URL", ""),
	}
}
