-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users
(name, password, email)
VALUES
($1, $2, $3)
RETURNING *;

