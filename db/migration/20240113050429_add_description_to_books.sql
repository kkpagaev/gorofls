-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books DROP COLUMN description;
-- +goose StatementEnd
