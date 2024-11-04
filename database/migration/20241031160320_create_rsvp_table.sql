-- migrate:up
CREATE TABLE rsvp (
    id bigserial NOT NULL,
    name varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL,
    guest_amount varchar(255) NOT NULL,
    is_attend INTEGER NOT NULL DEFAULT 0 CHECK (is_attend IN (0, 1)),
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT rsvp_pkey PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS rsvp