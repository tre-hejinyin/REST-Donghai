create database "myGolang"
    with owner postgres;

create user golang with password 'golang'
    superuser
    createdb
    createrole;

ALTER DATABASE "myGolang" SET timezone TO 'Asia/Shanghai';

create table if not exists users
(
    id         bigserial,
    first_name text,
    last_name  text,
    email      text not null
        constraint users_pk
            primary key,
    password   text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table users
    owner to golang;

create unique index if not exists users_email_uindex
    on users (email);
