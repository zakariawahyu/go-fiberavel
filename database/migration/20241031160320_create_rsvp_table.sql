-- migrate:up
CREATE TABLE rsvp (
    id bigserial NOT NULL,
    name varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL,
    guest_amount varchar(255) NOT NULL,
    is_attend bool NOT NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    CONSTRAINT rsvp_pkey PRIMARY KEY (id)
)

-- migrate:down
DROP TABLE IF EXISTS rsvp