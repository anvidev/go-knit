package handler

import (
	"go-starter/internal/service"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Landing *LandingHandler
	Auth    *AuthHandler
	Project *ProjectHandler
}

func New(s *service.Service) *Handler {
	return &Handler{
		Landing: NewLandingHandler(),
		Auth:    NewAuthHandler(s.User),
		Project: NewProjectHandler(),
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
