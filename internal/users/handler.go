package users

import (
	"go-fiber/pkg/templadapter"
	"go-fiber/pkg/validator"
	"go-fiber/views/components"
	"net/http"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type UsersHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *UserRepository
	store      *session.Store
}

func NewHandler(router fiber.Router, logger *zerolog.Logger, repository *UserRepository, store *session.Store) {
	h := &UsersHandler{
		router:     router,
		logger:     logger,
		repository: repository,
		store:      store,
	}
	h.router.Post("/registration", h.addUser)
	h.router.Post("/login", h.checkUser)
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
	component := components.Notification("Аккаунт успешно зарегистрирован", components.NotificationSuccess)
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *UsersHandler) checkUser(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
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
	)

	if len(errors.Errors) > 0 {
		component := components.Notification(validator.ParseErrors(*errors), components.NotificationFail)
		return templadapter.Render(c, component, http.StatusBadRequest)
	}
	user, err := h.repository.checkUser(form)
	if err != nil {
		if err.Error() == "Incorrect password" {
			h.logger.Error().Msg(err.Error())
			component := components.Notification("Неверный пароль", components.NotificationFail)
			return templadapter.Render(c, component, http.StatusBadRequest)
		}
		h.logger.Error().Msg(err.Error())
		component := components.Notification("Ошибка на сервере", components.NotificationFail)
		return templadapter.Render(c, component, http.StatusBadRequest)
	}

	// Создать сессию
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		panic(err)
	}

	// TODO - Перенаправить на главную
	// c.Redirect("/")
	// return nil
	component := components.Notification("Вы вошли в систему", components.NotificationSuccess)
	return templadapter.Render(c, component, http.StatusOK)
}
