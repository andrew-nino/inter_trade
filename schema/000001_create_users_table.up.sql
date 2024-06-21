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
    type_hash       varchar(10)  not null,
    key             varchar(255) not null,
    hash            char(64)     not null unique,
    created_at      timestamp    not null default now()
);