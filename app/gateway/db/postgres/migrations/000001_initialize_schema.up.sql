BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS transfers
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_origin_id text NOT NULL,
    account_destination_id text NOT NULL,
    amount int NOT NULL,
    created_at	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS accounts
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" text NOT NULL,
    cpf text NOT NULL,
    "secret" text NOT NULL,
    balance int NULL,
    created_at	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
