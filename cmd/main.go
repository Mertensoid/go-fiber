package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/internal/users"
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/database"
	"go-fiber/pkg/logger"
	"go-fiber/pkg/middleware"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	// Repositories
	vacancyRepository := vacancy.NewVacancyRepository(dbpool, customLogger)
	usersRepository := users.NewUserRepository(dbpool, customLogger)

	// Handlers
	home.NewHandler(app, customLogger, vacancyRepository, store)
	vacancy.NewHandler(app, customLogger, vacancyRepository)
	users.NewHandler(app, customLogger, usersRepository, store)
	app.Listen(":5001")
}
