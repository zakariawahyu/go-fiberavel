-- migrate:up
CREATE TABLE wishes (
   id bigserial NOT NULL,
   name varchar(255) NOT NULL,
   wish_description text NOT NULL,
   created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
   deleted_at timestamp(0) NULL,
   CONSTRAINT wishes_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS wishes