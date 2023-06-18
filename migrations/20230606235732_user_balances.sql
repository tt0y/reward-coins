-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_balances
(
    id           bigint unsigned auto_increment primary key,
    user_id      int,
    coin_type_id int,
    amount       int
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_balances;
-- +goose StatementEnd
