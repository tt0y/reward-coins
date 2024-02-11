-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, email, phone, password, salt, is_admin, activated_at, blocked_at, blocked_to, created_at, updated_at, deleted_at) VALUES ('Super Admin', 'sa@rc.com', '+77786373432', '5517baada0de60466bf6cc4b5e08ef37', 'qMMG6EFa', 1, '2024-02-11 13:54:18', NULL, NULL, '2024-02-11 13:54:18', '2024-02-11 13:54:18', NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email = 'sa@rc.com';
-- +goose StatementEnd
