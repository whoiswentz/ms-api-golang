CREATE
DATABASE bank;
SET
DATABASE = bank;

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
    amount       DECIMAL   NOT NULL,
    status       INT       NOT NULL DEFAULT 1,
    INDEX (customer_id)
);

INSERT INTO accounts
    (customer_id, opening_date, account_type, pin, amount, status)
VALUES ('4038c462-18e6-4841-9992-40bf50314ceb', '2020-08-22 10:20:06', 'Saving', '1075', 1000.0, 1),
       ('41f9d6c0-c3f3-4ecf-8956-81e259a9b7bb', '2020-03-12 8:12:03', 'Saving', '1076', 2000.0, 1),
       ('7dd96b28-dc65-48ab-a7cf-499130e0c1eb', '2021-01-12 1:12:03', 'Saving', '1077', 5000.0, 1);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions
(
    transaction_id   UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    account_id       UUID      NOT NULL REFERENCES accounts (account_id),
    amount           DECIMAL   NOT NULL,
    transaction_type STRING    NOT NULL,
    transaction_date TIMESTAMP NOT NULL DEFAULT current_timestamp(),
    INDEX (account_id)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    username    STRING PRIMARY KEY NOT NULL,
    password    STRING             NOT NULL,
    role        STRING             NOT NULL,
    customer_id UUID                        DEFAULT NULL,
    created_on  TIMESTAMP          NOT NULL DEFAULT current_timestamp()
);

INSERT INTO users
VALUES ('admin', 'abc123', 'admin', NULL, '2020-08-09 10:27:22'),
       ('2001', 'abc123', 'user', '4038c462-18e6-4841-9992-40bf50314ceb', '2020-08-09 10:27:22'),
       ('2000', 'abc123', 'user', '41f9d6c0-c3f3-4ecf-8956-81e259a9b7bb', '2020-08-09 10:27:22');

SELECT username, u.customer_id, role, array_agg(a.account_id) as account_numbers
FROM users u
         LEFT JOIN accounts a ON a.customer_id = u.customer_id
WHERE username = '2000'
  and password = 'abc123'
GROUP BY a.customer_id, username;