-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions
(
    id                  bigint unsigned auto_increment primary key,
    status              VARCHAR(30),
    transaction_type_id int,
    coin_type_id        int,
    amount              int,
    user_id_from        int,
    user_id_to          int,
    product_id          int,
    date                timestamp default CURRENT_TIMESTAMP
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
