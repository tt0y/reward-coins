-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions
(
    id                  bigint unsigned auto_increment primary key,
    status              VARCHAR(30)    default 'new',
    transaction_type_id int,
    coin_type_id        int,
    amount              int,
    user_id_from        int,
    user_id_to          int       null,
    product_id          int       null,
    created_at          timestamp      default CURRENT_TIMESTAMP,
    updated_at          timestamp      default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at          timestamp null default null
) charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
