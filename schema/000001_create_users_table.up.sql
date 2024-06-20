CREATE TABLE IF NOT EXISTS users 
(
    id serial not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    created_at timestamp not null default now()
);
-- CREATE TABLE IF NOT EXISTS cars
-- (
--     id            serial not null unique,
--     regnum        varchar(255) not null,
--     mark          varchar(255) not null,
--     model         varchar(255) not null,
--     year          int
-- );
-- CREATE TABLE IF NOT EXISTS owners
-- (
--     id            serial not null unique,
--     car_id        int references cars (id) on delete cascade not null,
--     name          varchar(255) not null,
--     surname       varchar(255) not null,
--     patronymic    varchar(255)
-- );