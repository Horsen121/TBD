BEGIN;

CREATE TABLE productList(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    time DATE,
    owner VARCHAR(255)
);

CREATE TABLE buyList(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    weight float,
    reminder DATE,
    owner VARCHAR(255)
);

CREATE TABLE lastProduct(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    owner VARCHAR(255),
    status bool,
    date DATE
);

CREATE TABLE users(
    id VARCHAR(255),
    name VARCHAR(255)
);

COMMIT;