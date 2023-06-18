-- +goose Up
-- +goose StatementBegin
CREATE TABLE transaction_types
(
    id         bigint unsigned auto_increment primary key,
    name       VARCHAR(20),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at timestamp null default null
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transaction_types;
-- +goose StatementEnd
