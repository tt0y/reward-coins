-- +goose Up
-- +goose StatementBegin
CREATE TABLE products
(
    id           bigint unsigned auto_increment primary key,
    name         VARCHAR(60),
    description  VARCHAR(255),
    cost         int,
    coin_type_id int,
    images       json      null,
    active       boolean        default true,
    created_at   timestamp      default CURRENT_TIMESTAMP,
    updated_at   timestamp      default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at   timestamp null default null
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
