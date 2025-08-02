package home

import (
	"bytes"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: *customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	tmpl, err := template.New("test").Parse("{{.Count}} -  количество пользователей")
	data := struct{ Count int }{Count: 1}
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template error")
	}
	var finalTemplateBuffer bytes.Buffer
	if err = tmpl.Execute(&finalTemplateBuffer, data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template compile error")
	}
	return c.Send(finalTemplateBuffer.Bytes())
}
func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().Msg("Hello")
	return c.SendString("Error")
}
