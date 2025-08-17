package home

import (
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/templadapter"
	"go-fiber/views"
	"go-fiber/views/pages"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router     fiber.Router
	logger     zerolog.Logger
	repository vacancy.VacancyRepository
}

type User struct {
	Id   int
	Name string
}

type Category struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, vacancy *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router:     router,
		logger:     *customLogger,
		repository: *vacancy,
	}
	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)
	h.router.Get("/login", h.login)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
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
	return templadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Info().Msg("Hello")
	return c.SendString("Error")
}
