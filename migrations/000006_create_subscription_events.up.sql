CREATE TABLE subscription_events (
    id UUID PRIMARY KEY,
    subscription_id UUID NOT NULL REFERENCES subscriptions(id),
    event_type subscription_event_type NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX subscription_events_subscription_id_idx
    ON subscription_events(subscription_id);
