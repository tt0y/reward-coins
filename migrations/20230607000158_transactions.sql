-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions
(
    id                  int NOT NULL,
    status              VARCHAR(30),
    transaction_type_id int,
    coin_type_id        int,
    amount              int,
    user_id_from        int,
    user_id_to          int,
    product_id          int,
    date                timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (transaction_type_id) REFERENCES transaction_types (id),
    FOREIGN KEY (coin_type_id) REFERENCES coin_types (id),
    FOREIGN KEY (user_id_from) REFERENCES users (id),
    FOREIGN KEY (user_id_to) REFERENCES users (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
