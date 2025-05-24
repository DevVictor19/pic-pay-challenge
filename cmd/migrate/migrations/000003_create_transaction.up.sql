DROP TYPE IF EXISTS transaction_type;
CREATE TYPE transaction_type AS ENUM ('payment_received', 'payment_sent');

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    payer_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    payee_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type transaction_type NOT NULL,
    amount BIGINT NOT NULL CHECK (amount > 0),
    description TEXT,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
