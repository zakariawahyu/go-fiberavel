-- migrate:up
CREATE TABLE guest_lists (
    id bigserial NOT NULL,
    name varchar(255) NOT NULL,
    slug varchar(255) NOT NULL,
    is_gift bool NOT NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    deleted_at timestamp(0) NULL,
    CONSTRAINT guest_lists_pkey PRIMARY KEY (id),
    CONSTRAINT guest_lists_slug_unique UNIQUE (slug)
);

-- migrate:down
DROP TABLE IF EXISTS guest_lists