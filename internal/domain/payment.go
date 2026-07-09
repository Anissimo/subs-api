package domain

import "time"

type PaymentStatus string

const (
	PaymentPending  PaymentStatus = "PENDING"
	PaymentSuccess  PaymentStatus = "SUCCESS"
	PaymentFailed   PaymentStatus = "FAILED"
	PaymentRefunded PaymentStatus = "REFUNDED"
)

type Payment struct {
	ID                    string
	SubscriptionID        string
	Amount                int64
	Currency              string
	Status                PaymentStatus
	ProviderTransactionID string
	PaidAt                *time.Time
	CreatedAt             time.Time
}
