
-- TODO: Create table based on Kubera features
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