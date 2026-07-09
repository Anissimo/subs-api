CREATE TABLE subscriptions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    plan_id UUID NOT NULL REFERENCES plans(id),
    status subscription_status NOT NULL,
    started_at TIMESTAMPTZ NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    auto_renew BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (expires_at > started_at)
);

CREATE UNIQUE INDEX subscriptions_one_active_per_user_idx
    ON subscriptions(user_id)
    WHERE status = 'ACTIVE';

CREATE INDEX subscriptions_user_id_idx ON subscriptions(user_id);
CREATE INDEX subscriptions_plan_id_idx ON subscriptions(plan_id);
CREATE INDEX subscriptions_expires_at_idx ON subscriptions(expires_at);
