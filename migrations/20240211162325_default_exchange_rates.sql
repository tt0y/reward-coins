-- +goose Up
-- +goose StatementBegin
insert into exchange_rates (coin_type_id_from, coin_type_id_to, rate, created_at, updated_at, deleted_at) values (1, 2, 1, '2024-02-11 16:23:25', '2024-02-11 16:23:25', null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from exchange_rates where coin_type_id_from = 1 and coin_type_id_to = 2;
-- +goose StatementEnd
