package handler

import (
	"go-starter/internal/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h HomeHandler) ShowHome(c echo.Context) error {
	return render(c, home.Show("Welcome to Knit"))
}
