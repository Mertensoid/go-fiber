package home

import (
	"go-fiber/pkg/templadapter"
	"go-fiber/views"

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
	h.router.Get("/", h.home)
	// h.router.Get("/cats", h.categories)
	h.router.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return templadapter.Render(c, component)
}

// func (h *HomeHandler) categories(c *fiber.Ctx) error {
// 	cats := []Category{
// 		{Id: 1, Name: "Еда"},
// 		{Id: 2, Name: "Животные"},
// 		{Id: 3, Name: "Машины"},
// 		{Id: 4, Name: "Спорт"},
// 		{Id: 5, Name: "Музыка"},
// 		{Id: 6, Name: "Технологии"},
// 		{Id: 7, Name: "Прочее"},
// 	}
// 	return c.Render("categories", cats)
// }

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().Msg("Hello")
	return c.SendString("Error")
}
