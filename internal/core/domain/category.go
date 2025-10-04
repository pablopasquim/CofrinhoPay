package domain

import "time"

type Category struct {
	ID        uint
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
