CREATE TYPE subscription_status AS ENUM (
    'ACTIVE',
    'PAUSED',
    'CANCELLED',
    'EXPIRED'
);

CREATE TYPE payment_status AS ENUM (
    'PENDING',
    'SUCCESS',
    'FAILED',
    'REFUNDED'
);

CREATE TYPE subscription_event_type AS ENUM (
    'CREATED',
    'RENEWED',
    'PLAN_CHANGED',
    'PAUSED',
    'RESUMED',
    'CANCELLED',
    'EXPIRED'
);
