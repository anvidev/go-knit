package handler

import (
	"fmt"
	"go-starter/internal/model"
	"go-starter/internal/service"
	"go-starter/internal/view/project"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService service.ProjectService
}

func NewProjectHandler(s service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: s,
	}
}

func (h ProjectHandler) ShowProjects(c echo.Context) error {
	return render(c, project.ShowProjects("Your knitting projects"))
}

func (h ProjectHandler) CreateProject(c echo.Context) error {
	return render(c, project.CreateProject("Create new project"))
}

func (h ProjectHandler) PostCreateProject(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	difficulty := "easy"
	userID := 1

	fmt.Println("THIS IS THE TITLE AND DESC", title, description)

	p := model.Project{
		Title:       title,
		Description: description,
		Difficulty:  difficulty,
		UserID:      userID,
	}

	id, err := h.projectService.Create(&p)
	if err != nil {
		return render(c, project.CreateForm(project.CreateFormData{
			Message: err.Error(),
		}))
	}

	return hxRedirect(c, fmt.Sprintf("/projects/%d", id))
}
