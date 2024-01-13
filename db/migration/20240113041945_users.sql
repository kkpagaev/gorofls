-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id       BIGSERIAL PRIMARY KEY,
  name     text      NOT NULL,
  emai     text      NOT NULL,
  password text  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
