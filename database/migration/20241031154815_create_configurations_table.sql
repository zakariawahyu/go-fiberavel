-- migrate:up
CREATE TABLE configurations (
    id bigserial NOT NULL,
    type varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    image varchar(255),
    image_caption varchar(255) NOT NULL,
    custom_data json NULL,
    is_active bool NOT NULL,
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT configurations_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS configurations