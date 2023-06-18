-- +goose Up
-- +goose StatementBegin
CREATE TABLE coin_types
(
    id          bigint unsigned auto_increment primary key,
    name        VARCHAR(60),
    description VARCHAR(255),
    created_at  timestamp default CURRENT_TIMESTAMP,
    updated_at  timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at  timestamp null default null
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coin_types;
-- +goose StatementEnd
