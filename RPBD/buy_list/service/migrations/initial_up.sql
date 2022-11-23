BEGIN;
CREATE TABLE users(
    id VARCHAR(255)
);

CREATE TABLE productList(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    time timestamp,
    FOREIGN KEY(owner)  REFERENCES users(id)
);

CREATE TABLE buyList(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    weight float
    reminder timestamp,
    FOREIGN KEY(owner)  REFERENCES users(id)
);

CREATE TABLE lastProduct(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    FOREIGN KEY(owner)  REFERENCES users(id),
    status bool,
    date timestamp
);

-- INSERT INTO people(name) VALUES
-- ('Владимир'),('Владислав'),('Даниил');

COMMIT;