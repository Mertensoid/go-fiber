package vacancy

import (
	"go-fiber/pkg/templadapter"
	"go-fiber/pkg/validator"
	"go-fiber/views/components"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	time.Sleep(time.Second * 2)
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или не верный",
		},
	)

	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.ParseErrors(*errors), components.NotificationFail)
	} else {
		component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	}
	return templadapter.Render(c, component)
}
