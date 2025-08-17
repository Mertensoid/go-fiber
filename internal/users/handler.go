package users

import (
	"go-fiber/pkg/templadapter"
	"go-fiber/pkg/validator"
	"go-fiber/views/components"
	"net/http"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type UsersHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *UserRepository
}

func NewHandler(router fiber.Router, logger *zerolog.Logger, repository *UserRepository) {
	h := &UsersHandler{
		router:     router,
		logger:     logger,
		repository: repository,
	}
	h.router.Post("/registration", h.addUser)
}

func (h *UsersHandler) addUser(c *fiber.Ctx) error {
	form := RegistrationForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Name:     c.FormValue("name"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или не верный",
		},
		&validators.StringIsPresent{
			Name:    "Password",
			Field:   form.Password,
			Message: "Пароль не задан",
		},
		&validators.StringIsPresent{
			Name:    "Name",
			Field:   form.Name,
			Message: "Имя пользователя не задано",
		},
	)

	if len(errors.Errors) > 0 {
		component := components.Notification(validator.ParseErrors(*errors), components.NotificationFail)
		return templadapter.Render(c, component, http.StatusBadRequest)
	}
	err := h.repository.addUser(form)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		component := components.Notification("Ошибка на сервере", components.NotificationFail)
		return templadapter.Render(c, component, http.StatusBadRequest)
	}
	component := components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return templadapter.Render(c, component, http.StatusOK)
}
