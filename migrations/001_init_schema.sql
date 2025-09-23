-- identity service

CREATE TABLE identity_users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE identity_user_identities (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES identity_users(id) ON DELETE CASCADE, 
    provider TEXT NOT NULL,
    external_id TEXT NOT NULL,
    secret_hash TEXT NOT NULL,
    meta JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    UNIQUE (provider, external_id)
);

CREATE TABLE identity_refresh_token (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES identity_users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL DEFAULT now(),
    revoked_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- purchase

CREATE TABLE purchase ( 
    id UUID PRIMARY KEY,
    owner_user_id UUID REFERENCES identity_users(id) ON DELETE SET NULL,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE purchase_invites (
    id UUID PRIMARY KEY,
    purchase_id UUID NOT NULL REFERENCES purchase(id) ON DELETE CASCADE,
    to_user_reference TEXT NOT NULL, -- reference to the user in external source(mob. number/email/tag)
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE purchase_tasks (
    id UUID PRIMARY KEY,
    purchase_id UUID NOT NULL REFERENCES purchase(id) ON DELETE CASCADE,
    author_user_id UUID REFERENCES identity_users(id) ON DELETE SET NULL,
    title TEXT NOT NULL,
    description TEXT,
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    assignee_user_id UUID REFERENCES identity_users(id) ON DELETE SET NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE purchase_participants (
    purchase_id UUID NOT NULL REFERENCES purchase(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES identity_users(id) ON DELETE CASCADE,
    display_name TEXT NOT NULL,
    joined_at TIMESTAMP NOT NULL DEFAULT now(),
    
    PRIMARY KEY (purchase_id, user_id)
);

-- payments

CREATE TABLE payments_intents (
    id UUID PRIMARY KEY,
    purchase_id UUID NOT NULL REFERENCES purchase(id),
    payer_user_id UUID NOT NULL REFERENCES identity_users(id),
    payee_user_id UUID NOT NULL REFERENCES identity_users(id),
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    status TEXT NOT NULL,
    due_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE payment_attempt (
    id UUID PRIMARY KEY,
    intent_id UUID REFERENCES payments_intents(id) ON DELETE SET NULL,
    provider TEXT NOT NULL,
    provider_reference TEXT NOT NULL, -- unique id of opertation in the provider database
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    status TEXT NOT NULL,
    confirmation_url TEXT,
    error JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE payments_refund (
    id UUID PRIMARY KEY,
    intent_id UUID REFERENCES payments_intents(id) ON DELETE SET NULL,
    attempt_id UUID REFERENCES payment_attempt(id) ON DELETE SET NULL,
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    status TEXT NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE payments_ledger_entries (
    id UUID PRIMARY KEY,
    occurred_at TIMESTAMP NOT NULL DEFAULT now(),
    account TEXT NOT NULL,
    side TEXT NOT NULL CHECK (side IN ('debit', 'credit')),
    amount_cents BIGINT NOT NULL,
    currency CHAR(3) NOT NULL,
    intent_id UUID REFERENCES payments_intents(id) ON DELETE SET NULL,
    attempt_id UUID REFERENCES payment_attempt(id) ON DELETE SET NULL
);