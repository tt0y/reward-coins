-- +goose Up
-- +goose StatementBegin
CREATE TABLE coin_types
(
    id          int NOT NULL,
    name        VARCHAR(60),
    description VARCHAR(255),
    created_at  timestamp,
    updated_at  timestamp,
    deleted_at  timestamp,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coin_types;
-- +goose StatementEnd
