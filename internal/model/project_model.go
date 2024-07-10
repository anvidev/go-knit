package model

type Project struct {
	ID          int64 `bun:"id,pk,autoincrement"`
	Title       string
	Description string
	Designer    string
	Yarn        string
	// Color       []string
	NeedleSize int
	// AmountUsed float32
	Difficulty string
	// Washing    string

	// StartDate time.Time
	// EndDate   time.Time

	Public bool
	UserID int
}
