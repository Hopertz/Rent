CREATE EXTENSION citext;

CREATE TABLE IF NOT EXISTS users (
    admin_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1
);


