BEGIN;

CREATE TABLE productList(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    time timestamp,
    owner VARCHAR(255)
);

CREATE TABLE buyList(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    weight float,
    reminder timestamp,
    owner VARCHAR(255)
);

CREATE TABLE lastProduct(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    owner VARCHAR(255),
    status bool,
    date timestamp
);

COMMIT;