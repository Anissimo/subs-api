package domain

import "time"

type SubscriptionEventType string

const (
	EventCreated     SubscriptionEventType = "CREATED"
	EventRenewed     SubscriptionEventType = "RENEWED"
	EventPlanChanged SubscriptionEventType = "PLAN_CHANGED"
	EventPaused      SubscriptionEventType = "PAUSED"
	EventResumed     SubscriptionEventType = "RESUMED"
	EventCancelled   SubscriptionEventType = "CANCELLED"
	EventExpired     SubscriptionEventType = "EXPIRED"
)

type SubscriptionEvent struct {
	ID             string
	SubscriptionID string
	EventType      SubscriptionEventType
	Description    string
	CreatedAt      time.Time
}
