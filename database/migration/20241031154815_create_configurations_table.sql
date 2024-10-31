-- migrate:up
CREATE TABLE configurations (
    id bigserial NOT NULL,
    type varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    image varchar(255) NULL,
    image_caption varchar(255) NULL,
    custom_data json NULL,
    is_active bool NOT NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    CONSTRAINT configurations_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS configurations