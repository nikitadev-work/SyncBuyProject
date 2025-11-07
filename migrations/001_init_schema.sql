-- identity service

CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT,
    status TEXT NOT NULL CHECK (status in ('active', 'blocked', 'deleted')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT right_name CHECK (first_name <> '')
);

CREATE TABLE user_identities (
    external_id TEXT NOT NULL,
    internal_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider_type TEXT NOT NULL CHECK (provider_type in ('telegram')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    chat_id TEXT NOT NULL,
    meta JSONB,

    UNIQUE (provider_type, external_id)
);

-- purchase

CREATE TABLE purchase ( 
    id UUID PRIMARY KEY,
    owner_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL CHECK (status in ('active', 'finished', 'locked')),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE purchase_invites (
    id UUID PRIMARY KEY,
    purchase_id UUID NOT NULL REFERENCES purchase(id) ON DELETE CASCADE,
    to_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status TEXT NOT NULL CHECK (status in ('active', 'rejected', 'accepted')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    UNIQUE (purchase_id, to_user_id)
);

CREATE TABLE purchase_tasks (
    id UUID PRIMARY KEY,
    purchase_id UUID NOT NULL REFERENCES purchase(id) ON DELETE CASCADE,
    author_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    title TEXT NOT NULL,
    description TEXT,
    amount_cents BIGINT NOT NULL CHECK (amount_cents >= 0),
    currency CHAR(3) NOT NULL,
    assignee_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE purchase_participants (
    purchase_id UUID NOT NULL REFERENCES purchase(id) ON DELETE CASCADE,
    participant_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    
    PRIMARY KEY (purchase_id, participant_id)
);

-- payments

CREATE TABLE payment_intents (
    id UUID PRIMARY KEY,
    purchase_id UUID NOT NULL REFERENCES purchase(id),
    payer_user_id UUID NOT NULL REFERENCES users(id),
    payee_user_id UUID NOT NULL REFERENCES users(id),
    amount_cents BIGINT NOT NULL CHECK (amount_cents >= 0),
    currency CHAR(3) NOT NULL,
    status TEXT NOT NULL CHECK (status in ('active', 'finished', 'canceled')),
    due_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE payment_attempt (
    id UUID PRIMARY KEY,
    intent_id UUID REFERENCES payment_intents(id),
    provider TEXT NOT NULL CHECK (provider in ('ukassa')),
    provider_payment_id TEXT NOT NULL, -- id of paymend in a provider system
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    status TEXT NOT NULL,
    confirmation_url TEXT,
    error JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    UNIQUE (provider, provider_payment_id)
);

CREATE TABLE payment_refund (
    id UUID PRIMARY KEY,
    intent_id UUID REFERENCES payment_intents(id),
    attempt_id UUID NOT NULL REFERENCES payment_attempt(id),
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    status TEXT NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE payment_ledger_entries (
    id UUID PRIMARY KEY,
    occurred_at TIMESTAMP NOT NULL DEFAULT now(),
    account TEXT NOT NULL,
    side TEXT NOT NULL CHECK (side IN ('debit', 'credit')),
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    intent_id UUID REFERENCES payment_intents(id) ON DELETE SET NULL,
    attempt_id UUID REFERENCES payment_attempt(id) ON DELETE SET NULL
);

CREATE INDEX identities_index ON user_identities (internal_id);