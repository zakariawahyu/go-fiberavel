-- migrate:up
CREATE TABLE admins (
    id bigserial NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT admins_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS admins