CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    phone_number text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    name text COLLATE pg_catalog."default",
    created_at date,
    updated_at date,
    deleted_at date
)