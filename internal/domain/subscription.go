package domain

import "time"

type SubscriptionStatus string

const (
	SubscriptionActive    SubscriptionStatus = "ACTIVE"
	SubscriptionPaused    SubscriptionStatus = "PAUSED"
	SubscriptionCancelled SubscriptionStatus = "CANCELLED"
	SubscriptionExpired   SubscriptionStatus = "EXPIRED"
)

type Subscription struct {
	ID        string
	UserID    string
	PlanID    string
	Status    SubscriptionStatus
	StartedAt time.Time
	ExpiresAt time.Time
	AutoRenew bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
