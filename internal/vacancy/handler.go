package vacancy

import (
	"go-fiber/pkg/templadapter"
	"go-fiber/pkg/validator"
	"go-fiber/views/components"
	"strconv"
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
	salary, _ := strconv.ParseInt(c.FormValue("salary"), 10, 64)
	form := VacancyCreateForm{
		Role:     c.FormValue("role"),
		Company:  c.FormValue("company"),
		Sphere:   c.FormValue("sphere"),
		Salary:   int(salary),
		Location: c.FormValue("location"),
		Email:    c.FormValue("email"),
	}
	time.Sleep(time.Second * 2)
	errors := validate.Validate(
		&validators.StringIsPresent{
			Name:    "Role",
			Field:   form.Role,
			Message: "Должность не задана",
		},
		&validators.StringIsPresent{
			Name:    "Company",
			Field:   form.Company,
			Message: "Название компании не задано",
		},
		&validators.StringIsPresent{
			Name:    "Sphere",
			Field:   form.Sphere,
			Message: "Сфера деятельности компании не задана",
		},
		&validators.IntIsGreaterThan{
			Name:     "Salary",
			Field:    form.Salary,
			Compared: 0,
			Message:  "Зароботная плата не задана",
		},
		&validators.StringIsPresent{
			Name:    "Location",
			Field:   form.Location,
			Message: "Расположение места работы не задано",
		},
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
