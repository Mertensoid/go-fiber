package vacancy

import (
	"go-fiber/pkg/templadapter"
	"go-fiber/views/components"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router fiber.Router
	logger *zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) {
	h := &VacancyHandler{
		router: router,
		logger: logger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	h.logger.Info().Msg(email)
	component := components.Notification("Вакансия успешно создана")
	return templadapter.Render(c, component)
}
