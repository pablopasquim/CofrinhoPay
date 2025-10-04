package domain

import "time"

type Account struct {
	ID        uint
	Name      string
	Type      string
	Balance   float64
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
