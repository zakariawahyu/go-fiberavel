-- migrate:up
CREATE TABLE public.gifts (
    id bigserial NOT NULL,
    bank varchar(255) NOT NULL,
    account_name varchar(255) NOT NULL,
    account_number varchar(255) NOT NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    deleted_at timestamp(0) NULL,
    CONSTRAINT gifts_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS gifts