package domain

import "time"

type Plan struct {
	ID           string
	Name         string
	Description  string
	Price        int64
	DurationDays int
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
