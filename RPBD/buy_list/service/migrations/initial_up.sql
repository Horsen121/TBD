BEGIN;

CREATE TABLE product_list(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    time DATE,
    owner VARCHAR(255)
);

CREATE TABLE buy_list(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    weight float,
    reminder DATE,
    owner VARCHAR(255)
);

CREATE TABLE last_product(
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