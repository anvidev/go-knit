package handler

import (
	"go-starter/internal/view/landing"

	"github.com/labstack/echo/v4"
)

type LandingHandler struct{}

func NewLandingHandler() *LandingHandler {
	return &LandingHandler{}
}

func (h LandingHandler) ShowLanding(c echo.Context) error {
	return render(c, landing.Show("Welcome to Knit"))
}
