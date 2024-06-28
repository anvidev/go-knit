package handler

import (
	"go-starter/internal/view/project"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {}

func NewProjectHandler() *ProjectHandler {
  return &ProjectHandler{}
}

func (h ProjectHandler) ShowProjects(c echo.Context) error {
  return render(c, project.ShowProjects("Your knitting projects"))
}
