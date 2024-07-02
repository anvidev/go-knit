package model

import "time"

type Project struct {
	ID         int `bun:"id,pk,autoincrement"`
	Title      string
	Content    string
	Designer   string
	Yarn       string
	Color      []string
	NeedleSize int
	AmountUsed float32
	Difficulty string
	Washing    string

	StartDate  time.Time
	EndDate    time.Time

	Public     bool
	UserID     int
}
