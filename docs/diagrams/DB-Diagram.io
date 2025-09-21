// IDENTITY / AUTH

Table identity_users {
  id uuid [pk, note: 'Global User ID']
  name text [not null]
  status text [not null, default: 'active']          // active|blocked|deleted
  created_at timestamptz [not null, default: `now()`]
}

Table identity_user_identities {
  id uuid [pk]
  user_id uuid [not null, ref: > identity_users.id]
  provider text [not null]                            // telegram|email|phone|google|apple|password
  external_id text [not null]                         // tg id / email / phone / oauth sub
  secret_hash text                                    // only for provider=password
  meta jsonb [not null, default: '{}']
  created_at timestamptz [not null, default: `now()`]
}

Table identity_refresh_tokens {
  id uuid [pk]
  user_id uuid [not null, ref: > identity_users.id]
  token_hash text [not null]
  expires_at timestamptz [not null]
  revoked_at timestamptz
  created_at timestamptz [not null, default: `now()`]
}

// PURCHASE

Table purchases {
  id uuid [pk]
  owner_user_id uuid [not null, ref: > identity_users.id]
  title text [not null]
  description text
  status text [not null, default: 'draft']            // draft|active|locked|closed
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz
}

Table purchase_participants {
  purchase_id uuid [not null, ref: > purchases.id]
  user_id uuid [not null, ref: > identity_users.id]
  display_name text [not null]
  joined_at timestamptz [not null, default: `now()`]
}

Table purchase_invites {
  id uuid [pk]
  purchase_id uuid [not null, ref: > purchases.id]
  to_user_reference text [not null]                   // email/phone/username/external id
  status text [not null, default: 'pending']          // pending|accepted|declined
  created_at timestamptz [not null, default: `now()`]
}

Table purchase_tasks {
  id uuid [pk]
  purchase_id uuid [not null, ref: > purchases.id]
  author_user_id uuid [not null, ref: > identity_users.id]
  title text [not null]
  description text
  amount_cents int [not null]
  currency text [not null, default: 'RUB']
  assignee_user_id uuid [ref: > identity_users.id]                       // nullable
  status text [not null, default: 'open']             // open|assigned|done|deleted
  created_at timestamptz [not null, default: `now()`]
}

// PAYMENTS

Table payments_intents {
  id uuid [pk]
  purchase_id uuid [not null, ref: > purchases.id]
  payer_user_id uuid [not null, ref: > identity_users.id]
  payee_user_id uuid [not null, ref: > identity_users.id]
  amount_cents int [not null]
  currency text [not null, default: 'RUB']
  status text [not null, default: 'created']          // created|awaiting|succeeded|canceled|expired|refunded
  due_at timestamptz
  created_at timestamptz [not null, default: `now()`]
}

Table payments_attempts {
  id uuid [pk]
  intent_id uuid [not null, ref: > payments_intents.id]
  provider text [not null]                            // yookassa|...
  provider_reference text                              // id at provider
  amount_cents int [not null]
  currency text [not null]
  status text [not null, default: 'created']          // created|pending_confirmation|succeeded|canceled|failed
  confirmation_url text
  error jsonb
  created_at timestamptz [not null, default: `now()`]
}

Table payments_refunds {
  id uuid [pk]
  intent_id uuid [ref: > payments_intents.id]
  attempt_id uuid [ref: > payments_attempts.id]
  amount_cents int [not null]
  currency text [not null]
  status text [not null, default: 'created']          // created|succeeded|failed
  reason text
  created_at timestamptz [not null, default: `now()`]
}

Table payments_ledger_entries {
  entry_id uuid [pk]
  occurred_at timestamptz [not null, default: `now()`]
  account text [not null]
  side text [not null]                                // debit|credit
  amount_cents int [not null]
  currency text [not null]
  purchase_id uuid
  intent_id uuid
  attempt_id uuid
}
