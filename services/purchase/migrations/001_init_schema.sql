-- purchase service

CREATE TABLE purchases (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    currency TEXT NOT NULL,
    settlement_initiated_at TIMESTAMP DEFAULT NULL,
    locked_at TIMESTAMP,
    finished_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT right_title CHECK (title <> ''),
    CONSTRAINT right_currency CHECK (currency in ('RUB'))
);

CREATE TABLE participants (
    user_id UUID NOT NULL,
    purchase_id UUID NOT NULL REFERENCES purchases(id) ON DELETE CASCADE,
    joined_at TIMESTAMP NOT NULL DEFAULT now(),

    PRIMARY KEY (user_id, purchase_id)
);

CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    purchase_id UUID NOT NULL REFERENCES purchases(id) ON DELETE CASCADE,
    author_user_id UUID NOT NULL,
    executor_user_id UUID,
    done BOOLEAN DEFAULT false,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT right_title CHECK (title <> ''),
    CONSTRAINT right_amount CHECK (amount > 0)
);
