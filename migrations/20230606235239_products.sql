-- +goose Up
-- +goose StatementBegin
CREATE TABLE products
(
    id           bigint unsigned auto_increment primary key,
    name         VARCHAR(255),
    cost         int,
    coin_type_id int,
    active       boolean default true
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
