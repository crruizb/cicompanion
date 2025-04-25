CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username varchar NOT NULL UNIQUE,
    github_pat varchar NOT NULL
);
