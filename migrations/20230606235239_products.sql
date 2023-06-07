-- +goose Up
-- +goose StatementBegin
CREATE TABLE products
(
    id           int NOT NULL,
    name         VARCHAR(255),
    cost         int,
    coin_type_id int,
    active       boolean,
    pictures     json,
    created_at   timestamp,
    updated_at   timestamp,
    deleted_at   timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (coin_type_id) REFERENCES coin_types (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
