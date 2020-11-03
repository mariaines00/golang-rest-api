CREATE DATABASE factory;

CREATE USER asimov PASSWORD 'password';

GRANT ALL PRIVILEGES ON DATABASE factory TO asimov;

CREATE TABLE robots (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(50) NOT NULL UNIQUE,
    model INT NOT NULL
);

CREATE TABLE friendships (
    id SERIAL PRIMARY KEY,
    buddy1 INT NOT NULL REFERENCES robots(id),
    buddy2 INT NOT NULL REFERENCES robots(id)
);
