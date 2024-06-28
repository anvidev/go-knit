package service

import "go-starter/internal/database"

type Service struct {
  User UserService
}

func New(db *database.Database) *Service {
  return &Service{
    User: NewUserService(db),
  }
}
