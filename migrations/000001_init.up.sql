CREATE TABLE authors
(
    id      serial       not null unique,
    name    varchar(255) not null,
    surname varchar(255) not null,
    country varchar(255) not null
);

CREATE TABLE books
(
    id         serial                                        not null unique,
    title      varchar(255)                                  not null,
    author_id  int references authors (id) on DELETE cascade not null,
    impression varchar(255)                                  not null
);

CREATE TABLE users
(
    id       serial       not null unique,
    username varchar(255) not null,
    password varchar(255) not null
);