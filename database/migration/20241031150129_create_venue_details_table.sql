-- migrate:up
CREATE TABLE venue_details (
      id bigserial NOT NULL,
      name varchar(255) NOT NULL,
      location varchar(255) NOT NULL,
      address varchar(255) NOT NULL,
      date timestamp(0) NOT NULL,
      map text NOT NULL,
      created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
      deleted_at timestamp(0) NULL,
      CONSTRAINT venue_details_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS venue_details