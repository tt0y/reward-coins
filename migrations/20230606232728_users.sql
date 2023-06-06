-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id           int NOT NULL,
    name         varchar(255),
    email        varchar(255),
    phone        varchar(30),
    password     varchar(60),
    is_admin     boolean,
    activated_at timestamp,
    blocked_at   timestamp,
    blocked_to   timestamp,
    created_at   timestamp,
    updated_at   timestamp,
    deleted_at   timestamp,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
