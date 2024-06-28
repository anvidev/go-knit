package service

import (
	"context"
	"go-starter/internal/database"
	"go-starter/internal/model"
	"time"
)

type UserService interface {
	Create(string, string, string) (*model.User, error)
	GetByEmail(string) (*model.User, error)
}

type userService struct {
	db *database.Database
}

func NewUserService(db *database.Database) UserService {
	return &userService{db: db}
}

func (s userService) Create(name, email, hashedPassword string) (*model.User, error) {
	now := time.Now()
	newUser := &model.User{
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	_, err := s.db.Postgres.NewInsert().Model(newUser).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s userService) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := s.db.Postgres.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background()); err != nil {
		return nil, err
	}
	return &user, nil
}
