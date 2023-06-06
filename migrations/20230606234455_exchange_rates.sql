-- +goose Up
-- +goose StatementBegin
CREATE TABLE exchange_rates
(
    id                int NOT NULL,
    coin_type_id_from int,
    coin_type_id_to   int,
    rate              int,
    created_at        timestamp,
    updated_at        timestamp,
    deleted_at        timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (coin_type_id_from) REFERENCES coin_types (id),
    FOREIGN KEY (coin_type_id_to) REFERENCES coin_types (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE exchange_rates;
-- +goose StatementEnd
