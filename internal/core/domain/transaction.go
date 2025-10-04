package domain

import "time"

type Transaction struct {
	ID         uint
	UserID     uint
	AccountID  uint
	CategoryID uint
	Amount     float64
	Type       string
	Date       time.Time
	Note       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
