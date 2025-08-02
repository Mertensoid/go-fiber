package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger zerolog.Logger
}

type User struct {
	Id   int
	Name string
}

type Category struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: *customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/cats", h.categories)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	names := []string{"Anton", "Vasya"}
	users := []User{
		{Id: 1, Name: "Anton"},
		{Id: 2, Name: "Vasya"},
	}
	data := struct {
		Count   int
		IsAdmin bool
		CanUse  bool
		IsRange bool
		Names   []string
		Users   []User
	}{Count: 1, IsAdmin: true, CanUse: true, IsRange: false, Names: names, Users: users}
	if data.IsRange {
		return c.Render("range_page", data)
	}
	return c.Render("page", data)
}

func (h *HomeHandler) categories(c *fiber.Ctx) error {
	cats := []Category{
		{Id: 1, Name: "Еда"},
		{Id: 1, Name: "Животные"},
		{Id: 1, Name: "Машины"},
		{Id: 1, Name: "Спорт"},
		{Id: 1, Name: "Музыка"},
		{Id: 1, Name: "Технологии"},
		{Id: 1, Name: "Прочее"},
	}
	return c.Render("categories", cats)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().Msg("Hello")
	return c.SendString("Error")
}
