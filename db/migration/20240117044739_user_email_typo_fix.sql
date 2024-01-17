-- +goose Up
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN emai TO email;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN email TO emai;
-- +goose StatementEnd
