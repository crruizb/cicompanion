CREATE TABLE IF NOT EXISTS monitoring_urls(
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL,
    url varchar NOT NULL,
    status varchar CHECK (status IN ('OK', 'ERROR')),
    user_id uuid NOT NULL REFERENCES users(id)
);

ALTER TABLE users ADD role varchar NOT NULL DEFAULT USER;