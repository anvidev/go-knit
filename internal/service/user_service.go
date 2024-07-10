package service

import (
	"context"
	"go-starter/internal/model"
	"go-starter/internal/store"
	"time"
)

type UserService interface {
	Create(string, string, string) (*model.User, error)
	GetByEmail(string) (*model.User, error)
}

type userService struct {
	store *store.Store
}

func NewUserService(store *store.Store) UserService {
	return &userService{store: store}
}

func (s userService) Create(name, email, hashedPassword string) (*model.User, error) {
	now := time.Now()
	newUser := &model.User{
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
		CreatedAt:      now,
	}
	_, err := s.store.DB.NewInsert().Model(newUser).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s userService) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := s.store.DB.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background()); err != nil {
		return nil, err
	}
	return &user, nil
}
