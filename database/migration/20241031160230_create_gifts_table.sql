-- migrate:up
CREATE TABLE gifts (
    id bigserial NOT NULL,
    bank varchar(255) NOT NULL,
    account_name varchar(255) NOT NULL,
    account_number varchar(255) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(0) NULL,
    CONSTRAINT gifts_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS gifts