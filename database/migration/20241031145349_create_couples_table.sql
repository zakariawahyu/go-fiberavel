-- migrate:up
CREATE TABLE couples (
    id bigserial NOT NULL,
    couple_type varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    parent_description varchar(255) NOT NULL,
    father_name varchar(255) NOT NULL,
    mother_name varchar(255) NOT NULL,
    image varchar(255) NOT NULL,
    image_caption varchar(255) NOT NULL,
    instagram_url varchar(255) NOT NULL,
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(0) NULL,
    CONSTRAINT couples_couple_type_check CHECK (((couple_type)::text = ANY ((ARRAY['cpp'::character varying, 'cpw'::character varying])::text[]))),
    CONSTRAINT couples_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS couples