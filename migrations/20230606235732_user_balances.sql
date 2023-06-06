-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_balances
(
    id           int NOT NULL,
    user_id      int,
    coin_type_id int,
    amount       int,
    PRIMARY KEY (id),
    FOREIGN KEY (coin_type_id) REFERENCES coin_types (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_balances;
-- +goose StatementEnd
