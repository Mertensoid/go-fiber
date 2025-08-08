package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/database"
	"go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbConfig := config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")

	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	home.NewHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger)
	app.Listen(":5001")
}
