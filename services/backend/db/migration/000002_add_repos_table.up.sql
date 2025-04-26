CREATE TABLE repos (
    id int NOT NULL,
    name varchar NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id),
    PRIMARY KEY (id, user_id)
);
