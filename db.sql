CREATE
DATABASE banking;
SET
DATABASE = banking;

DROP TABLE IF EXISTS customer;
CREATE TABLE customers
(
    customer_id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name          STRING NOT NULL,
    date_of_birth DATE   NOT NULL,
    city          STRING NOT NULL,
    zipcode       STRING NOT NULL,
    status        INT    NOT NULL  DEFAULT 1
);

INSERT INTO customers
    (name, date_of_birth, city, zipcode, status)
VALUES ('Steve', '1978-12-15', 'Delhi', '666', 1),
       ('Arien', '1988-12-21', 'NY', '666', 1),
       ('QWEQ', '1931-12-21', 'DC', '666', 1);

DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts
(
    account_id   UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    customer_id  UUID      NOT NULL REFERENCES customers (customer_id),
    opening_date TIMESTAMP NOT NULL DEFAULT current_timestamp(),
    account_type STRING    NOT NULL,
    pin          STRING    NOT NULL,
    status       INT       NOT NULL DEFAULT 1,
    INDEX (customer_id)
);

INSERT INTO accounts
    (customer_id, opening_date, account_type, pin, status)
VALUES ('4d75ae30-cb09-4339-95c9-be48e080afb8', '2020-08-22 10:20:06', 'Saving', '1075', 1),
       ('65c911fa-9528-4716-b728-53df228bd6b0', '2020-03-12 8:12:03', 'Saving', '1076', 1),
       ('a155508b-881f-4a00-aa28-4a861d0d48ee', '2021-01-12 1:12:03', 'Saving', '1077', 1);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts (account_id),
    amount DECIMAL NOT NULL,
    transaction_type STRING NOT NULL,
    transaction_date TIMESTAMP NOT NULL DEFAULT current_timestamp(),
    INDEX (account_id)
);