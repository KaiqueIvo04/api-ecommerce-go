CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table users (
    id uuid primary key default uuid_generate_v4(),
    name varchar(255) not null,
    email varchar(30) not null,
    password varchar(20) not null,
    type varchar(15) not null,
    created_at timestamp not null default now()
);