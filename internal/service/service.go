package service

import "go-starter/internal/store"

type Service struct {
	User    UserService
	Project ProjectService
}

func New(store *store.Store) *Service {
	return &Service{
		User:    NewUserService(store),
		Project: NewProjectService(store),
	}
}
