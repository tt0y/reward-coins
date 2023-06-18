-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id           bigint unsigned auto_increment primary key,
    name         varchar(255),
    email        varchar(255),
    phone        varchar(30),
    password     varchar(60),
    is_admin     boolean,
    activated_at timestamp,
    blocked_at   timestamp,
    blocked_to   timestamp,
    created_at   timestamp default CURRENT_TIMESTAMP,
    updated_at   timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at   timestamp null default null
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
