package service

import (
	"context"
	"go-starter/internal/model"
	"go-starter/internal/store"
)

type ProjectService interface {
	GetAll() ([]model.Project, error)
	GetOne(int64) (*model.Project, error)
	Create(*model.Project) (int64, error)
}

type projectService struct {
	store *store.Store
}

func NewProjectService(s *store.Store) ProjectService {
	return &projectService{
		store: s,
	}
}

func (s projectService) Create(p *model.Project) (int64, error) {
	_, err := s.store.DB.NewInsert().Model(p).Returning("id").Exec(context.Background())
	if err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (s projectService) GetOne(id int64) (*model.Project, error) {
	var p model.Project
	if err := s.store.DB.NewSelect().Model(&p).Where("id = ?", id).Scan(context.TODO()); err != nil {
		return nil, err
	}
	return &p, nil
}

func (s projectService) GetAll() ([]model.Project, error) {
	var projects []model.Project
	if err := s.store.DB.NewSelect().Model(&projects).Scan(context.Background()); err != nil {
		return nil, err
	}
	return projects, nil
}
