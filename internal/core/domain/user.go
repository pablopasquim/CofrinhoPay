package domain

import "time"

type User struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
