-- migrate:up
CREATE TABLE wishes (
   id bigserial NOT NULL,
   name varchar(255) NOT NULL,
   wish_description text NOT NULL,
   created_at timestamp(0) NULL,
   updated_at timestamp(0) NULL,
   deleted_at timestamp(0) NULL,
   CONSTRAINT wishes_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS wishes