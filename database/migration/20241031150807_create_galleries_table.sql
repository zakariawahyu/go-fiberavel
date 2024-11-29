-- migrate:up
CREATE TABLE galleries (
    id bigserial NOT NULL,
    image varchar(255),
    image_caption varchar(255) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(0) NULL,
    CONSTRAINT galleries_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS galleries