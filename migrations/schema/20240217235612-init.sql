
-- +migrate Up

create table if not exists users (
    id         uuid        primary key not null default gen_random_uuid(),
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz default null,
    username   text        not null unique,
    email      text        not null unique,
    password   text        not null,
    otp_secret text        default null
);

-- +migrate Down
