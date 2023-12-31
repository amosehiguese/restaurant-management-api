CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS menu(
    id              UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name            VARCHAR (200) NOT NULL,
    description     TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS dish(
    id              UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name            VARCHAR (200) NOT NULL,
    description     TEXT NOT NULL DEFAULT '',
    price           DECIMAL (20, 2) NOT NULL,
    menu_id         UUID NOT NULL REFERENCES menu (id) ON DELETE CASCADE
);

CREATE TYPE order_status AS ENUM ('pending', 'processing', 'completed', 'canceled');

CREATE TABLE IF NOT EXISTS orders(
    id                 UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    status             order_status     NOT NULL,
    created_at         TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP NULL   

);

CREATE TABLE IF NOT EXISTS orderItem(
    id             UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    order_id       UUID NOT NULL REFERENCES orders (id)   ON DELETE CASCADE,
    dish_id        UUID NOT NULL REFERENCES dish (id) ON DELETE CASCADE,
    quantity       INT NOT NULL

);

CREATE TYPE restaurant_table_status AS ENUM ('occupied', 'available');

CREATE TABLE IF NOT EXISTS restaurant_table(
    id              UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    number          INT  NOT NULL,
    capacity        INT  NOT NULL,
    status          restaurant_table_status   NOT NULL

);

CREATE TYPE reservation_status AS ENUM ('available', 'confirmed', 'canceled');

CREATE TABLE IF NOT EXISTS reservation(
    id                     UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    table_id               UUID NOT NULL REFERENCES restaurant_table (id) ON DELETE CASCADE,
    reservation_date       DATE NOT NULL,
    reservation_time       TIME  NOT NULL, 
    status                 reservation_status  NOT NULL,
    created_at             TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMP NULL   
);

CREATE TABLE IF NOT EXISTS invoice(
    id                 UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    order_id           UUID NOT NULL REFERENCES orders (id)   ON DELETE CASCADE,
    invoice_date       DATE NOT NULL,
    total_amount       DECIMAL (20, 2) NOT NULL,
    tax                DECIMAL (20, 2) NOT NULL,
    discount           DECIMAL (20, 2) NOT NULL,
    grand_total        DECIMAL (20, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS roles(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(50) NOT NULL UNIQUE,
    description     TEXT NOT NULL DEFAULT ''
);


CREATE TABLE IF NOT EXISTS users(
    id                 UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    first_name         VARCHAR(50) NOT NULL,
    last_name          VARCHAR(50) NOT NULL ,
    username           VARCHAR(255) NOT NULL UNIQUE,
    email              VARCHAR(255) NOT NULL UNIQUE,
    password_hash      VARCHAR(255) NOT NULL,
    user_role          INT NOT NULL REFERENCES roles (id) ON DELETE CASCADE, 
    created_at         TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP NULL   
);



