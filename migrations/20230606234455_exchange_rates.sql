-- +goose Up
-- +goose StatementBegin
CREATE TABLE exchange_rates
(
    id                bigint unsigned auto_increment primary key,
    coin_type_id_from int,
    coin_type_id_to   int,
    rate              int,
    created_at        timestamp      default CURRENT_TIMESTAMP,
    updated_at        timestamp      default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at        timestamp null default null
)
    charset = utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE exchange_rates;
-- +goose StatementEnd
