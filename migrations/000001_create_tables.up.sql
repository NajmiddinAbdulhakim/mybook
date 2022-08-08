CREATE TABLE IF NOT EXISTS books (
    id uuid NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    photo_link VARCHAR(255),
    author_name uuid NOT NULL
);

CREATE TABLE IF NOT EXISTS authors (
    id uuid NOT NULL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    photo_link VARCHAR(255) 
);

CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255)
);