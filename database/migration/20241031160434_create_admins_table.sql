-- migrate:up
CREATE TABLE admins (
    id bigserial NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    CONSTRAINT admins_pkey PRIMARY KEY (id)
);

-- migrate:down
DRROP TABLE IF EXISTS admins