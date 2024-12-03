-- migrate:up
CREATE TABLE configurations (
    id bigserial NOT NULL,
    type varchar(255) UNIQUE NOT NULL,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    image varchar(255),
    image_caption varchar(255),
    custom_data jsonb,
    is_active bool DEFAULT true,
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT configurations_pkey PRIMARY KEY (id),
    CONSTRAINT unique_type UNIQUE (type)
);

-- migrate:down
DROP TABLE IF EXISTS configurations