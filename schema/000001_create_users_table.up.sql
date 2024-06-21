CREATE TABLE IF NOT EXISTS users
(
    id              serial       not null unique,
    username        varchar(255) not null unique,
    password_hash   varchar(255) not null,
    created_at      timestamp    not null default now()
);

CREATE TABLE IF NOT EXISTS hash_storage
(
    id              serial       not null unique,
    key             varchar(255) not null,
    hash            bytea        not null unique,
    created_at      timestamp    not null default now()
);