package repository

import "context"

type UserRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type PlanRepository interface{}

type SubscriptionRepository interface{}

type PaymentRepository interface{}

type SubscriptionEventRepository interface{}
