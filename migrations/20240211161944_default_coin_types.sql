-- +goose Up
-- +goose StatementBegin
INSERT INTO coin_types (name, description, created_at, updated_at, deleted_at)
VALUES
('green coin', 'Basic green coin. It is credited at the beginning of the month and burns at the end of the month if not used.', '2024-02-11 16:13:54', '2024-02-11 16:13:54', null),
('red coin', 'Basic red coin, into which green coins are converted when transferred from user to user.', '2024-02-11 16:13:54', '2024-02-11 16:13:54', null),
('extra coin', 'Extra coin. Given by Higher Powers.', '2024-02-11 16:13:54', '2024-02-11 16:13:54', null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM coin_types WHERE name IN ('green coin', 'red coin', 'extra coin');
-- +goose StatementEnd
