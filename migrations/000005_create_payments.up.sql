CREATE TABLE payments (
    id UUID PRIMARY KEY,
    subscription_id UUID NOT NULL REFERENCES subscriptions(id),
    amount BIGINT NOT NULL CHECK (amount >= 0),
    currency CHAR(3) NOT NULL,
    status payment_status NOT NULL,
    provider_transaction_id TEXT,
    paid_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX payments_subscription_id_idx ON payments(subscription_id);
CREATE UNIQUE INDEX payments_provider_transaction_id_idx
    ON payments(provider_transaction_id)
    WHERE provider_transaction_id IS NOT NULL;
