-- +goose Up
-- +goose StatementBegin
INSERT INTO user_balances (user_id, coin_type_id, amount, created_at, updated_at, deleted_at) VALUES
(1, 1, 100000000, '2024-02-11 16:23:44', '2024-02-11 16:23:44', null), -- green coin. SU gives these coins monthly
(1, 3, 1000, '2024-02-11 16:23:44', '2024-02-11 16:23:44', null); -- SU gives these coins for special merits
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from user_balances where user_id = 1 and coin_type_id in (1,3);
-- +goose StatementEnd
