package home

import (
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/templadapter"
	"go-fiber/views"
	"go-fiber/views/pages"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *vacancy.VacancyRepository
	store      *session.Store
}

type User struct {
	Id   int
	Name string
}

type Category struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, vacancy *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:     router,
		logger:     customLogger,
		repository: vacancy,
		store:      store,
	}
	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)
	h.router.Get("/login", h.login)
	h.router.Get("/registration", h.registration)
	h.router.Get("/logout", h.logout)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)

	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	userEmail := ""
	if email, ok := session.Get("email").(string); ok {
		userEmail = email
	}
	c.Locals("email", userEmail)

	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}

	component := views.Main(vacancies, int(math.Ceil(float64(count/PAGE_ITEMS))), page)
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := pages.Login()
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	userEmail := ""
	if email, ok := session.Get("email").(string); ok {
		userEmail = email
	}
	c.Locals("email", userEmail)
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) logout(c *fiber.Ctx) error {
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	session.Delete("email")
	if err := session.Save(); err != nil {
		panic(err)
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}

func (h *HomeHandler) registration(c *fiber.Ctx) error {
	component := pages.Registration()
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	userEmail := ""
	if email, ok := session.Get("email").(string); ok {
		userEmail = email
	}
	c.Locals("email", userEmail)
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Info().Msg("Hello")
	return c.SendString("Error")
}
