package model

import "time"

const SessionUserKey = "user"

type User struct {
	ID             int `bun:"id,pk,autoincrement"`
	Name           string
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Session struct {
	ID        int `bun:"id,pk,autoincrement"`
	UserID    int
	Token     string
	ExpiresAt time.Time
}

func (s Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

type SessionAttributes struct {
	Name     string
	Email    string
	LoggedIn bool
}
