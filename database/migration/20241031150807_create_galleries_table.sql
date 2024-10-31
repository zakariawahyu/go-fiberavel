-- migrate:up
CREATE TABLE public.galleries (
    id bigserial NOT NULL,
    image varchar(255) NOT NULL,
    image_caption varchar(255) NOT NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    deleted_at timestamp(0) NULL,
    CONSTRAINT galleries_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS galleries