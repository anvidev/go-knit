package handler

import (
	"go-starter/internal/service"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Home *HomeHandler
	Auth *AuthHandler
}

func New(s *service.Service) *Handler {
	return &Handler{
		Home: NewHomeHandler(),
		Auth: NewAuthHandler(s.User),
	}
}

func render(c echo.Context, comp templ.Component) error {
	return comp.Render(c.Request().Context(), c.Response())
}

func hxRedirect(c echo.Context, url string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", url)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}
	return c.Redirect(http.StatusSeeOther, url)
}

