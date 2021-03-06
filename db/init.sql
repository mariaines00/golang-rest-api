
CREATE USER asimov PASSWORD 'password';

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO asimov;
GRANT SELECT, UPDATE, USAGE ON ALL SEQUENCES IN SCHEMA public to asimov;

CREATE TABLE IF NOT EXISTS robots (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(50) NOT NULL UNIQUE,
    model INT NOT NULL
);

CREATE TABLE IF NOT EXISTS friendships (
    id SERIAL PRIMARY KEY,
    buddy1 INT NOT NULL REFERENCES robots(id),
    buddy2 INT NOT NULL REFERENCES robots(id)
);
