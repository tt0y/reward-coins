-- +goose Up
-- +goose StatementBegin
CREATE TABLE transaction_types
(
    id          int NOT NULL,
    name        VARCHAR(20),
    created_at  timestamp,
    updated_at  timestamp,
    deleted_at  timestamp,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transaction_types;
-- +goose StatementEnd
