CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT,
    address TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE asset_info (
    id SERIAL PRIMARY KEY,
    account_info JSON NOT NULL,
    institution_info JSON NOT NULL,
    asset_info JSON NOT NULL,
    sensible_data TEXT NOT NULL,
    descriptions TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE bank_account (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    asset_info_id SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE crypto_account (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    asset_info_id SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    asset_info_id SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE collectibles (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    asset_info_id SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    asset_info_id SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE private_shares (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    asset_info_id SERIAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);